package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/joshsoftware/peerly-backend/internal/pkg/apperrors"
	"github.com/joshsoftware/peerly-backend/internal/pkg/dto"
	"github.com/joshsoftware/peerly-backend/internal/repository"
	logger "github.com/sirupsen/logrus"
)

type coreValueStore struct {
	DB *sqlx.DB
}

func NewCoreValueRepo(db *sqlx.DB) repository.CoreValueStorer {
	return &coreValueStore{
		DB: db,
	}
}

func (cs *coreValueStore) ListCoreValues(ctx context.Context) (coreValues []repository.CoreValue, err error) {
	queryBuilder := sq.Select("id", "name", "description", "parent_core_value_id").From("core_values")
	listCoreValuesQuery, _, err := queryBuilder.ToSql()
	if err != nil {
		logger.Error(fmt.Sprintf("error in generating squirrel query, err: %s", err.Error()))
		return
	}
	err = cs.DB.SelectContext(
		ctx,
		&coreValues,
		listCoreValuesQuery,
	)

	if err != nil {
		logger.Error(fmt.Sprintf("error while getting core values, err: %s", err.Error()))
		return
	}

	return
}

func (cs *coreValueStore) GetCoreValue(ctx context.Context, coreValueID int64) (coreValue repository.CoreValue, err error) {
	queryBuilder := sq.
		Select("id", "name", "description", "parent_core_value_id").
		From("core_values").
		Where(sq.Eq{"id": coreValueID})

	getCoreValueQuery, args, err := queryBuilder.ToSql()
	if err != nil {
		logger.Error(fmt.Sprintf("error in generating squirrel query, err: %s", err.Error()))
		return
	}

	err = cs.DB.GetContext(
		ctx,
		&coreValue,
		getCoreValueQuery,
		args...,
	)
	if err != nil {
		logger.Error(fmt.Sprintf("error while getting core value, corevalue_id: %d, err: %s", coreValueID, err.Error()))
		err = apperrors.InvalidCoreValueData
		return
	}

	return
}

func (cs *coreValueStore) CreateCoreValue(ctx context.Context, coreValue dto.CreateCoreValueReq) (resp repository.CoreValue, err error) {

	queryBuilder := sq.Insert("core_values").Columns("name",
		"description", "parent_core_value_id").Values(coreValue.Name, coreValue.Description, coreValue.ParentCoreValueID).Suffix("RETURNING id, name, description, parent_core_value_id")

	createCoreValueQuery, args, err := queryBuilder.ToSql()
	if err != nil {
		logger.Error(fmt.Sprintf("error in generating squirrel query, err: %s", err.Error()))
		return
	}

	err = cs.DB.GetContext(
		ctx,
		&resp,
		createCoreValueQuery,
		args...,
	)
	if err != nil {
		logger.Error(fmt.Sprintf("error while creating core value, err: %s", err.Error()))
		return
	}

	return
}

func (cs *coreValueStore) UpdateCoreValue(ctx context.Context, updateReq dto.UpdateQueryRequest) (resp repository.CoreValue, err error) {
	queryBuilder := sq.Update("core_values").
		Set("name", updateReq.Name).
		Set("description", updateReq.Description).
		Where(sq.Eq{"id": updateReq.Id}).
		Suffix("RETURNING id, name, description, parent_core_value_id")

	updateCoreValueQuery, args, err := queryBuilder.ToSql()
	if err != nil {
		logger.Error(fmt.Sprintf("error in generating squirrel query, err: %s", err.Error()))
		return
	}
	err = cs.DB.GetContext(
		ctx,
		&resp,
		updateCoreValueQuery,
		args...,
	)
	if err != nil {
		logger.Error(fmt.Sprintf("error while updating core value, corevalue_id: %d, err: %s", updateReq.Id, err.Error()))
		return
	}

	return
}

func (cs *coreValueStore) CheckUniqueCoreVal(ctx context.Context, name string) (isUnique bool, err error) {

	isUnique = false
	resp := []int64{}

	queryBuilder := sq.Select("id").
		From("core_values").
		Where(sq.Eq{"name": name})

	checkUniqueCoreVal, args, err := queryBuilder.ToSql()
	if err != nil {
		logger.Error(fmt.Sprintf("error in generating squirrel query, err: %s", err.Error()))
		return
	}

	err = cs.DB.SelectContext(
		ctx,
		&resp,
		checkUniqueCoreVal,
		args...,
	)

	if err != nil {
		logger.Error(fmt.Sprintf("error while checking unique core vlaues, err: %s", err.Error()))
		err = apperrors.InternalServerError
		return
	}

	if len(resp) <= 0 {
		isUnique = true
		return
	}

	return
}
