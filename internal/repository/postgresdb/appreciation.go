package repository

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/joshsoftware/peerly-backend/internal/pkg/apperrors"
	"github.com/joshsoftware/peerly-backend/internal/pkg/constants"
	"github.com/joshsoftware/peerly-backend/internal/pkg/dto"
	"github.com/joshsoftware/peerly-backend/internal/repository"
	logger "github.com/sirupsen/logrus"
)

type appreciationsStore struct {
	BaseRepository
}

func NewAppreciationRepo(db *sqlx.DB) repository.AppreciationStorer {
	return &appreciationsStore{
		BaseRepository: BaseRepository{db},
	}
}

func (appr *appreciationsStore) CreateAppreciation(ctx context.Context, tx repository.Transaction, appreciation dto.Appreciation) (repository.Appreciation, error) {

	insertQuery, args, err := sq.
		Insert("appreciations").Columns(constants.CreateAppreciationColumns...).
		Values(appreciation.CoreValueID, appreciation.Description, appreciation.Quarter, appreciation.Sender, appreciation.Receiver).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING \"id\",\"core_value_id\", \"description\",\"total_reward_points\",\"quarter\",\"sender\",\"receiver\",\"created_at\",\"updated_at\"").
		ToSql()

	if err != nil {
		logger.Error(err.Error())
		return repository.Appreciation{}, apperrors.InternalServer
	}
	queryExecutor := appr.InitiateQueryExecutor(tx)
	var resAppr repository.Appreciation
	err = queryExecutor.QueryRowx(insertQuery, args...).Scan(&resAppr.ID, &resAppr.CoreValueID, &resAppr.Description, &resAppr.TotalRewards, &resAppr.Quarter, &resAppr.Sender, &resAppr.Receiver, &resAppr.CreatedAt, &resAppr.UpdatedAt)
	if err != nil {
		logger.Error("Error executing create certificate insert query: ", err)
		return repository.Appreciation{}, apperrors.InternalServer
	}

	return resAppr, nil
}

func (appr *appreciationsStore) GetAppreciationById(ctx context.Context, tx repository.Transaction, apprId int) (repository.AppreciationInfo, error) {

	// Build the SQL query
	query, args, err := sq.Select(
		"a.id",
		"cv.name AS core_value_name",
		"a.description",
		"a.is_valid",
		"a.total_reward_points",
		"a.quarter",
		"u_sender.first_name AS sender_first_name",
		"u_sender.last_name AS sender_last_name",
		"u_sender.profile_image_url AS sender_image_url",
		"u_sender.designation AS sender_designation",
		"u_receiver.first_name AS receiver_first_name",
		"u_receiver.last_name AS receiver_last_name",
		"u_receiver.profile_image_url AS receiver_image_url",
		"u_receiver.designation AS receiver_designation",
		"a.created_at",
		"a.updated_at",
	).From("appreciations a").
		LeftJoin("users u_sender ON a.sender = u_sender.id").
		LeftJoin("users u_receiver ON a.receiver = u_receiver.id").
		LeftJoin("core_values cv ON a.core_value_id = cv.id").
		PlaceholderFormat(sq.Dollar).
		Where(sq.And{
			sq.Eq{"a.id": apprId},
			sq.Eq{"is_valid": true},
		}).
		ToSql()

	if err != nil {
		logger.Error("err ", err.Error())
		return repository.AppreciationInfo{}, apperrors.InternalServer
	}

	queryExecutor := appr.InitiateQueryExecutor(tx)

	var resAppr repository.AppreciationInfo

	// Execute the query
	err = queryExecutor.QueryRowx(query, args...).Scan(
		&resAppr.ID,
		&resAppr.CoreValueName,
		&resAppr.Description,
		&resAppr.IsValid,
		&resAppr.TotalRewards,
		&resAppr.Quarter,
		&resAppr.SenderFirstName,
		&resAppr.SenderLastName,
		&resAppr.SenderImageURL,
		&resAppr.SenderDesignation,
		&resAppr.ReceiverFirstName,
		&resAppr.ReceiverLastName,
		&resAppr.ReceiverImageURL,
		&resAppr.ReceiverDesignation,
		&resAppr.CreatedAt,
		&resAppr.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Warn("no appreciation found with id: ", apprId)
			return repository.AppreciationInfo{}, apperrors.AppreciationNotFound
		}
		logger.Error("failed to execute query: ", err.Error())
		return repository.AppreciationInfo{}, apperrors.InternalServer
	}
	return resAppr, nil
}

func (appr *appreciationsStore) GetAppreciation(ctx context.Context, tx repository.Transaction, filter dto.AppreciationFilter, userID int64) ([]repository.AppreciationInfo, repository.Pagination, error) {

	// query builder for counting total records
	countQueryBuilder := sq.Select("COUNT(*)").
		From("appreciations a").
		LeftJoin("users u_sender ON a.sender = u_sender.id").
		LeftJoin("users u_receiver ON a.receiver = u_receiver.id").
		LeftJoin("core_values cv ON a.core_value_id = cv.id").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"a.is_valid": true})

	if filter.Name != "" {
		countQueryBuilder = countQueryBuilder.Where(
			"(CONCAT(u_sender.first_name, ' ', u_sender.last_name) LIKE ? OR "+
				"CONCAT(u_receiver.first_name, ' ', u_receiver.last_name) LIKE ?)",
			fmt.Sprintf("%%%s%%", filter.Name), fmt.Sprintf("%%%s%%", filter.Name),
		)
	}

	countSql, countArgs, err := countQueryBuilder.ToSql()
	if err != nil {
		logger.Error("failed to build count query: ", err.Error())
		return []repository.AppreciationInfo{}, repository.Pagination{}, apperrors.InternalServerError
	}

	queryExecutor := appr.InitiateQueryExecutor(tx)
	var totalRecords int64
	err = queryExecutor.QueryRowx(countSql, countArgs...).Scan(&totalRecords)
	if err != nil {
		logger.Error("failed to execute count query: ", err.Error())
		return []repository.AppreciationInfo{}, repository.Pagination{}, apperrors.InternalServerError
	}

	pagination := GetPaginationMetaData(filter.Page, filter.Limit, totalRecords)
	fmt.Println("pagination: ", pagination)

	// Initialize the Squirrel query builder
	queryBuilder := sq.Select(
		"a.id",
		"cv.name AS core_value_name",
		"cv.description AS core_value_description",
		"a.description",
		"a.is_valid",
		"a.total_reward_points",
		"a.quarter",
		"u_sender.first_name AS sender_first_name",
		"u_sender.last_name AS sender_last_name",
		"u_sender.profile_image_url AS sender_image_url",
		"u_sender.designation AS sender_designation",
		"u_receiver.first_name AS receiver_first_name",
		"u_receiver.last_name AS receiver_last_name",
		"u_receiver.profile_image_url AS receiver_image_url",
		"u_receiver.designation AS receiver_designation",
		"a.created_at",
		"a.updated_at",
		"COUNT(r.id) AS total_rewards",
		fmt.Sprintf(
			`COALESCE((
				SELECT r2.point 
				FROM rewards r2 
				WHERE r2.appreciation_id = a.id AND r2.sender = %d
			), 0) AS given_reward_point`, userID),
	).From("appreciations a").
		LeftJoin("users u_sender ON a.sender = u_sender.id").
		LeftJoin("users u_receiver ON a.receiver = u_receiver.id").
		LeftJoin("core_values cv ON a.core_value_id = cv.id").
		LeftJoin("rewards r ON a.id = r.appreciation_id").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"a.is_valid": true}).
		GroupBy("a.id, cv.name, cv.description, u_sender.first_name, u_sender.last_name, u_sender.profile_image_url, u_sender.designation, u_receiver.first_name, u_receiver.last_name, u_receiver.profile_image_url, u_receiver.designation")

	if filter.Name != "" {
		queryBuilder = queryBuilder.Where(
			"(CONCAT(u_sender.first_name, ' ', u_sender.last_name) LIKE ? OR "+
				"CONCAT(u_receiver.first_name, ' ', u_receiver.last_name) LIKE ?)",
			fmt.Sprintf("%%%s%%", filter.Name), fmt.Sprintf("%%%s%%", filter.Name),
		)
	}

	if filter.SortOrder != "" {
		queryBuilder = queryBuilder.OrderBy(fmt.Sprintf("a.created_at %s", filter.SortOrder))
	}

	offset := (filter.Page - 1) * filter.Limit

	// Add pagination
	queryBuilder = queryBuilder.Limit(uint64(filter.Limit)).Offset(uint64(offset))
	sql, args, err := queryBuilder.ToSql()
	if err != nil {
		logger.Error("failed to build query: ", err.Error())
		return nil, repository.Pagination{}, apperrors.InternalServerError
	}

	queryExecutor = appr.InitiateQueryExecutor(tx)
	res := make([]repository.AppreciationInfo, 0)
	err = sqlx.Select(queryExecutor, &res, sql, args...)
	if err != nil {
		logger.Error("failed to execute query: ", err.Error())
		return nil, repository.Pagination{}, apperrors.InternalServerError
	}

	return res, pagination, nil
}


func (appr *appreciationsStore) ValidateAppreciation(ctx context.Context, tx repository.Transaction, isValid bool, apprId int) (bool, error) {
	query, args, err := sq.Update("appreciations").
		Set("is_valid", isValid).
		Where(sq.And{
			sq.Eq{"id": apprId},
			sq.Eq{"is_valid": true},
		}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		logger.Error("Error building SQL: ", err.Error())
		return false, apperrors.InternalServer
	}

	queryExecutor := appr.InitiateQueryExecutor(tx)

	result, err := queryExecutor.Exec(query, args...)
	if err != nil {
		logger.Error("Error executing SQL: ", err.Error())
		return false, apperrors.InternalServer
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("Error getting rows affected: ", err.Error())
		return false, apperrors.InternalServer
	}

	if rowsAffected == 0 {
		logger.Warn("No rows affected")
		return false, apperrors.AppreciationNotFound
	}

	return true, nil
}

func (appr *appreciationsStore) IsUserPresent(ctx context.Context, tx repository.Transaction, userID int64) (bool, error) {
	// Initialize the Squirrel query builder
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	// Build the SQL query
	query, args, err := psql.Select("COUNT(*)").
		From("users").
		Where(sq.Eq{"id": userID}).
		ToSql()

	if err != nil {
		logger.Error("err ", err.Error())
		return false, apperrors.InternalServer
	}

	queryExecutor := appr.InitiateQueryExecutor(tx)

	var count int
	// Execute the query
	err = queryExecutor.QueryRowx(query, args...).Scan(&count)
	if err != nil {
		logger.Error("failed to execute query: ", err.Error())
		return false, apperrors.InternalServer
	}

	// Check if user is present
	return count > 0, nil
}
