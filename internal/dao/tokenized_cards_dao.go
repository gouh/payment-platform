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
	TokenizedCardDaoInterface interface {
		GetTokenizedCards(requests.TokenizedCardParams) ([]models.TokenizedCard, *int, error)
		GetTokenizedCardById(string, string) (*models.TokenizedCard, error)
		CreateTokenizedCard(models.TokenizedCard) (*models.TokenizedCard, error)
		DeleteTokenizedCard(tokenizedCardId string) error
	}
	TokenizedCardDao struct {
		Db        *pgxpool.Pool
		DbDialect *goqu.DialectWrapper
	}
)

func (dao *TokenizedCardDao) GetTokenizedCardById(customerId string, token string) (*models.TokenizedCard, error) {
	query, args, errBuildQuery := dao.buildTokenizedCardQuery(requests.TokenizedCardParams{Token: &token, CustomerId: &customerId}, false)
	if errBuildQuery != nil {
		return nil, errtrace.Wrap(errBuildQuery)
	}

	tokenizedCard := models.TokenizedCard{}
	ctx := context.Background()
	errGet := pgxscan.Get(ctx, dao.Db, &tokenizedCard, query, args...)

	if errGet != nil {
		if errGet.Error() == "scanning one: no rows in result set" {
			return nil, nil
		}
		return nil, errtrace.Wrap(errGet)
	}

	return &tokenizedCard, nil
}

func (dao *TokenizedCardDao) CreateTokenizedCard(tokenizedCard models.TokenizedCard) (*models.TokenizedCard, error) {
	ctx := context.Background()

	ds := goqu.Insert("tokenized_cards").
		Cols("token", "last_four_digits", "expiry_month", "expiry_year", "card_type", "customer_id").
		Vals(goqu.Vals{tokenizedCard.Token, tokenizedCard.LastFourDigits, tokenizedCard.ExpiryMonth, tokenizedCard.ExpiryYear, tokenizedCard.CardType, tokenizedCard.CustomerID}).
		Returning("token", "last_four_digits", "expiry_month", "expiry_year", "card_type", "customer_id")

	insertSQL, insertArgs, errInsertSql := ds.ToSQL()
	if errInsertSql != nil {
		return nil, errtrace.Wrap(errInsertSql)
	}

	errInsert := pgxscan.Get(ctx, dao.Db, &tokenizedCard, insertSQL, insertArgs...)
	if errInsert != nil {
		if strings.Contains(errInsert.Error(), "duplicate key value violates unique constraint") {
			return nil, nil
		}
		return nil, errtrace.Wrap(errInsert)
	}

	return &tokenizedCard, nil
}

func (dao *TokenizedCardDao) DeleteTokenizedCard(tokenizedCardId string) error {
	ctx := context.Background()

	ds := dao.DbDialect.Delete("tokenized_cards").Where(goqu.C("token").Eq(tokenizedCardId))
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

func (dao *TokenizedCardDao) GetTokenizedCards(params requests.TokenizedCardParams) ([]models.TokenizedCard, *int, error) {
	ctx := context.Background()

	var tokenizedCards []models.TokenizedCard
	query, args, errBuildQuery := dao.buildTokenizedCardQuery(params, false)
	if errBuildQuery != nil {
		return nil, nil, errtrace.Wrap(errBuildQuery)
	}
	errOnSelect := pgxscan.Select(ctx, dao.Db, &tokenizedCards, query, args...)
	if errOnSelect != nil {
		return nil, nil, errtrace.Wrap(errOnSelect)
	}

	var count int
	queryCount, _, errBuildQuery := dao.buildTokenizedCardQuery(params, true)
	if errBuildQuery != nil {
		return nil, nil, errtrace.Wrap(errBuildQuery)
	}
	errCount := dao.Db.QueryRow(ctx, queryCount).Scan(&count)
	if errCount != nil {
		return nil, nil, errtrace.Wrap(errCount)
	}

	return tokenizedCards, &count, nil
}

func (dao *TokenizedCardDao) buildTokenizedCardQuery(params requests.TokenizedCardParams, count bool) (string, []interface{}, error) {
	var ds *goqu.SelectDataset
	if count {
		ds = dao.DbDialect.From("tokenized_cards").Select(goqu.L("COUNT(*)").As("count"))
	} else {
		ds = dao.DbDialect.From("tokenized_cards").Select("*")
		ds = ds.Limit(uint(params.PageSize)).Offset(uint((params.Page - 1) * params.PageSize))
	}

	if params.CustomerId != nil {
		ds = ds.Where(goqu.Ex{"customer_id": params.CustomerId})
	}

	if params.Token != nil {
		ds = ds.Where(goqu.Ex{"token": params.Token})
	}

	return ds.ToSQL()
}

func NewTokenizedCardDao(container *container.Container) TokenizedCardDaoInterface {
	return &TokenizedCardDao{
		Db:        container.Db,
		DbDialect: container.DbDialect,
	}
}
