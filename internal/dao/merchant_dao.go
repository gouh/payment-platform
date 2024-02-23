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
	MerchantDaoInterface interface {
		GetMerchants(requests.MerchantParams) ([]models.Merchant, *int, error)
		GetMerchantById(string) (*models.Merchant, error)
		CreateMerchant(models.Merchant) (*models.Merchant, error)
		UpdateMerchant(merchant models.Merchant) (*models.Merchant, error)
		DeleteMerchant(merchantId string) error
	}
	MerchantDao struct {
		Db        *pgxpool.Pool
		DbDialect *goqu.DialectWrapper
	}
)

func (dao *MerchantDao) GetMerchantById(merchantId string) (*models.Merchant, error) {
	_, errUuid := uuid.Parse(merchantId)
	if errUuid != nil {
		return nil, errtrace.Wrap(errUuid)
	}

	query, args, errBuildQuery := dao.buildMerchantQuery(requests.MerchantParams{Id: &merchantId}, false)
	if errBuildQuery != nil {
		return nil, errtrace.Wrap(errBuildQuery)
	}

	merchant := models.Merchant{}
	ctx := context.Background()
	errGet := pgxscan.Get(ctx, dao.Db, &merchant, query, args...)

	if errGet != nil {
		if errGet.Error() == "scanning one: no rows in result set" {
			return nil, nil
		}
		return nil, errtrace.Wrap(errGet)
	}

	return &merchant, nil
}

func (dao *MerchantDao) CreateMerchant(merchant models.Merchant) (*models.Merchant, error) {
	ctx := context.Background()

	ds := goqu.Insert("merchants").
		Cols("name").
		Vals(goqu.Vals{merchant.Name}).
		Returning("id", "name")

	insertSQL, insertArgs, errInsertSql := ds.ToSQL()
	if errInsertSql != nil {
		return nil, errtrace.Wrap(errInsertSql)
	}

	errInsert := pgxscan.Get(ctx, dao.Db, &merchant, insertSQL, insertArgs...)
	if errInsert != nil {
		return nil, errtrace.Wrap(errInsert)
	}

	return &merchant, nil
}

func (dao *MerchantDao) UpdateMerchant(merchant models.Merchant) (*models.Merchant, error) {
	ctx := context.Background()

	ds := goqu.Update("merchants").Set(
		goqu.Record{"name": merchant.Name},
	).Where(goqu.C("id").Eq(merchant.ID)).Returning("*")

	updateSQL, updateArgs, updateErr := ds.ToSQL()
	if updateErr != nil {
		return nil, errtrace.Wrap(updateErr)
	}

	errInsert := pgxscan.Get(ctx, dao.Db, &merchant, updateSQL, updateArgs...)
	if errInsert != nil {
		return nil, errtrace.Wrap(errInsert)
	}

	return &merchant, nil
}

func (dao *MerchantDao) DeleteMerchant(merchantId string) error {
	ctx := context.Background()
	_, err := uuid.Parse(merchantId)
	if err != nil {
		return errtrace.Wrap(err)
	}

	ds := dao.DbDialect.Delete("merchants").Where(goqu.C("id").Eq(merchantId))
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

func (dao *MerchantDao) GetMerchants(params requests.MerchantParams) ([]models.Merchant, *int, error) {
	ctx := context.Background()

	var merchants []models.Merchant
	query, args, errBuildQuery := dao.buildMerchantQuery(params, false)
	if errBuildQuery != nil {
		return nil, nil, errtrace.Wrap(errBuildQuery)
	}
	errOnSelect := pgxscan.Select(ctx, dao.Db, &merchants, query, args...)
	if errOnSelect != nil {
		return nil, nil, errtrace.Wrap(errOnSelect)
	}

	var count int
	queryCount, _, errBuildQuery := dao.buildMerchantQuery(params, true)
	if errBuildQuery != nil {
		return nil, nil, errtrace.Wrap(errBuildQuery)
	}
	errCount := dao.Db.QueryRow(ctx, queryCount).Scan(&count)
	if errCount != nil {
		return nil, nil, errtrace.Wrap(errCount)
	}

	return merchants, &count, nil
}

func (dao *MerchantDao) buildMerchantQuery(params requests.MerchantParams, count bool) (string, []interface{}, error) {
	var ds *goqu.SelectDataset
	if count {
		ds = dao.DbDialect.From("merchants").Select(goqu.L("COUNT(*)").As("count"))
	} else {
		ds = dao.DbDialect.From("merchants").Select("*")
		ds = ds.Limit(uint(params.PageSize)).Offset(uint((params.Page - 1) * params.PageSize))
	}

	if params.Id != nil {
		ds = ds.Where(goqu.Ex{"id": params.Id})
	}

	return ds.ToSQL()
}

func NewMerchantDao(container *container.Container) MerchantDaoInterface {
	return &MerchantDao{
		Db:        container.Db,
		DbDialect: container.DbDialect,
	}
}
