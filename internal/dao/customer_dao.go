package dao

import (
	"braces.dev/errtrace"
	"context"
	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"payment-platform/internal/container"
	"payment-platform/internal/models"
	"payment-platform/internal/requests"
)

type (
	CustomerDaoInterface interface {
		GetCustomers(requests.CustomerParams) ([]models.Customer, *int, error)
		GetCustomerById(string) (*models.Customer, error)
		CreateCustomer(models.Customer) (*models.Customer, error)
		UpdateCustomer(customer models.Customer) (*models.Customer, error)
		DeleteCustomer(customerId string) error
	}
	CustomerDao struct {
		Db        *pgxpool.Pool
		DbDialect *goqu.DialectWrapper
	}
)

func (dao *CustomerDao) GetCustomerById(customerId string) (*models.Customer, error) {
	_, errUuid := uuid.Parse(customerId)
	if errUuid != nil {
		return nil, errtrace.Wrap(errUuid)
	}

	query, args, errBuildQuery := dao.buildCustomerQuery(requests.CustomerParams{Id: &customerId}, false)
	if errBuildQuery != nil {
		return nil, errtrace.Wrap(errBuildQuery)
	}

	customer := models.Customer{}
	ctx := context.Background()
	errGet := pgxscan.Get(ctx, dao.Db, &customer, query, args...)

	if errGet != nil {
		if errGet.Error() == "scanning one: no rows in result set" {
			return nil, nil
		}
		return nil, errtrace.Wrap(errGet)
	}

	return &customer, nil
}

func (dao *CustomerDao) CreateCustomer(customer models.Customer) (*models.Customer, error) {
	ctx := context.Background()

	ds := goqu.Insert("customers").
		Cols("name", "email").
		Vals(goqu.Vals{customer.Name, customer.Email}).
		Returning("id", "name", "email")

	insertSQL, insertArgs, errInsertSql := ds.ToSQL()
	if errInsertSql != nil {
		return nil, errtrace.Wrap(errInsertSql)
	}

	errInsert := pgxscan.Get(ctx, dao.Db, &customer, insertSQL, insertArgs...)
	if errInsert != nil {
		return nil, errtrace.Wrap(errInsert)
	}

	return &customer, nil
}

func (dao *CustomerDao) UpdateCustomer(customer models.Customer) (*models.Customer, error) {
	ctx := context.Background()

	ds := goqu.Update("customers").Set(
		goqu.Record{"name": customer.Name},
	).Where(goqu.C("id").Eq(customer.ID)).Returning("*")

	updateSQL, updateArgs, updateErr := ds.ToSQL()
	if updateErr != nil {
		return nil, errtrace.Wrap(updateErr)
	}

	errInsert := pgxscan.Get(ctx, dao.Db, &customer, updateSQL, updateArgs...)
	if errInsert != nil {
		return nil, errtrace.Wrap(errInsert)
	}

	return &customer, nil
}

func (dao *CustomerDao) DeleteCustomer(customerId string) error {
	ctx := context.Background()
	_, err := uuid.Parse(customerId)
	if err != nil {
		return errtrace.Wrap(err)
	}

	ds := dao.DbDialect.Delete("customers").Where(goqu.C("id").Eq(customerId))
	deleteSQL, deleteArgs, deleteErr := ds.ToSQL()
	if deleteErr != nil {
		return errtrace.Wrap(deleteErr)
	}

	cmdTag, err := dao.Db.Exec(ctx, deleteSQL, deleteArgs...)
	if err != nil {
		return errtrace.Wrap(err)
	}

	if cmdTag.RowsAffected() == 0 {
		return errtrace.Wrap(errtrace.New("no rows affected"))
	}

	return nil
}

func (dao *CustomerDao) GetCustomers(params requests.CustomerParams) ([]models.Customer, *int, error) {
	ctx := context.Background()

	var customers []models.Customer
	query, args, errBuildQuery := dao.buildCustomerQuery(params, false)
	if errBuildQuery != nil {
		return nil, nil, errtrace.Wrap(errBuildQuery)
	}
	errOnSelect := pgxscan.Select(ctx, dao.Db, &customers, query, args...)
	if errOnSelect != nil {
		return nil, nil, errtrace.Wrap(errOnSelect)
	}

	var count int
	queryCount, _, errBuildQuery := dao.buildCustomerQuery(params, true)
	if errBuildQuery != nil {
		return nil, nil, errtrace.Wrap(errBuildQuery)
	}
	errCount := dao.Db.QueryRow(ctx, queryCount).Scan(&count)
	if errCount != nil {
		return nil, nil, errtrace.Wrap(errCount)
	}

	return customers, &count, nil
}

func (dao *CustomerDao) buildCustomerQuery(params requests.CustomerParams, count bool) (string, []interface{}, error) {
	var ds *goqu.SelectDataset
	if count {
		ds = dao.DbDialect.From("customers").Select(goqu.L("COUNT(*)").As("count"))
	} else {
		ds = dao.DbDialect.From("customers").Select("*")
		ds = ds.Limit(uint(params.PageSize)).Offset(uint((params.Page - 1) * params.PageSize))
	}

	if params.Id != nil {
		ds = ds.Where(goqu.Ex{"id": params.Id})
	}

	if params.Name != nil {
		ds = ds.Where(goqu.C("name").Like("%" + *params.Name + "%"))
	}

	if params.Email != nil {
		ds = ds.Where(goqu.C("email").Like("%" + *params.Email + "%"))
	}

	return ds.ToSQL()
}

func NewCustomerDao(container *container.Container) CustomerDaoInterface {
	return &CustomerDao{
		Db:        container.Db,
		DbDialect: container.DbDialect,
	}
}
