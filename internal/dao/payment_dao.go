package dao

import (
	"braces.dev/errtrace"
	"context"
	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"payment-platform/internal/container"
	"payment-platform/internal/models"
	"payment-platform/internal/requests"
	"strings"
)

type (
	PaymentDaoInterface interface {
		GetPayments(requests.PaymentParams) ([]models.Payment, *int, error)
		GetPaymentById(string) (*models.Payment, error)
		CreatePayment(models.Payment) (*models.Payment, error)
	}
	PaymentDao struct {
		Db        *pgxpool.Pool
		DbDialect *goqu.DialectWrapper
	}
)

func (dao *PaymentDao) GetPaymentById(paymentId string) (*models.Payment, error) {
	query, args, errBuildQuery := dao.buildPaymentQuery(requests.PaymentParams{Id: &paymentId}, false)
	if errBuildQuery != nil {
		return nil, errtrace.Wrap(errBuildQuery)
	}

	payment := models.Payment{}
	ctx := context.Background()
	errGet := pgxscan.Get(ctx, dao.Db, &payment, query, args...)

	if errGet != nil {
		if errGet.Error() == "scanning one: no rows in result set" {
			return nil, nil
		}
		return nil, errtrace.Wrap(errGet)
	}

	return &payment, nil
}

func (dao *PaymentDao) CreatePayment(payment models.Payment) (*models.Payment, error) {
	ctx := context.Background()

	ds := goqu.Insert("payments").
		Cols("merchant_id", "token", "amount", "status", "transaction_date").
		Vals(goqu.Vals{payment.MerchantID, payment.Token, payment.Amount, payment.Status, payment.TransactionDate}).
		Returning("id", "merchant_id", "token", "amount", "status", "transaction_date")

	insertSQL, insertArgs, errInsertSql := ds.ToSQL()
	if errInsertSql != nil {
		return nil, errtrace.Wrap(errInsertSql)
	}

	errInsert := pgxscan.Get(ctx, dao.Db, &payment, insertSQL, insertArgs...)
	if errInsert != nil {
		if strings.Contains(errInsert.Error(), "duplicate key value violates unique constraint") {
			return nil, nil
		}
		return nil, errtrace.Wrap(errInsert)
	}

	return &payment, nil
}

func (dao *PaymentDao) GetPayments(params requests.PaymentParams) ([]models.Payment, *int, error) {
	ctx := context.Background()

	var payments []models.Payment
	query, args, errBuildQuery := dao.buildPaymentQuery(params, false)
	if errBuildQuery != nil {
		return nil, nil, errtrace.Wrap(errBuildQuery)
	}
	errOnSelect := pgxscan.Select(ctx, dao.Db, &payments, query, args...)
	if errOnSelect != nil {
		return nil, nil, errtrace.Wrap(errOnSelect)
	}

	var count int
	queryCount, _, errBuildQuery := dao.buildPaymentQuery(params, true)
	if errBuildQuery != nil {
		return nil, nil, errtrace.Wrap(errBuildQuery)
	}
	errCount := dao.Db.QueryRow(ctx, queryCount).Scan(&count)
	if errCount != nil {
		return nil, nil, errtrace.Wrap(errCount)
	}

	return payments, &count, nil
}

func (dao *PaymentDao) buildPaymentQuery(params requests.PaymentParams, count bool) (string, []interface{}, error) {
	var ds *goqu.SelectDataset
	if count {
		ds = dao.DbDialect.From("payments").Select(goqu.L("COUNT(*)").As("count")).
			Join(goqu.T("tokenized_cards"), goqu.On(goqu.Ex{"payments.token": goqu.I("tokenized_cards.token")})).
			Join(goqu.T("customers"), goqu.On(goqu.Ex{"tokenized_cards.customer_id": goqu.I("customers.id")}))
	} else {
		ds = dao.DbDialect.From("payments").Select("payments.*").
			Join(goqu.T("tokenized_cards"), goqu.On(goqu.Ex{"payments.token": goqu.I("tokenized_cards.token")})).
			Join(goqu.T("customers"), goqu.On(goqu.Ex{"tokenized_cards.customer_id": goqu.I("customers.id")}))
		ds = ds.Limit(uint(params.PageSize)).Offset(uint((params.Page - 1) * params.PageSize))
	}

	if params.Id != nil {
		ds = ds.Where(goqu.Ex{"payments.id": params.Id})
	}

	if params.CustomerId != nil {
		ds = ds.Where(goqu.Ex{"customers.id": params.CustomerId})
	}

	if params.TokenizedCard != nil {
		ds = ds.Where(goqu.Ex{"tokenized_cards.token": params.TokenizedCard})
	}

	if params.MerchantId != nil {
		ds = ds.Where(goqu.Ex{"payments.merchant_id": params.MerchantId})
	}

	return ds.ToSQL()
}

func NewPaymentDao(container *container.Container) PaymentDaoInterface {
	return &PaymentDao{
		Db:        container.Db,
		DbDialect: container.DbDialect,
	}
}
