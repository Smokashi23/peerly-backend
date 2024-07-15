package app

import (
	"github.com/jmoiron/sqlx"
	"github.com/joshsoftware/peerly-backend/internal/app/appreciation"
	corevalues "github.com/joshsoftware/peerly-backend/internal/app/coreValues"
	reward "github.com/joshsoftware/peerly-backend/internal/app/reward"
	reportappreciations "github.com/joshsoftware/peerly-backend/internal/app/reportAppreciations"

	user "github.com/joshsoftware/peerly-backend/internal/app/users"
	repository "github.com/joshsoftware/peerly-backend/internal/repository/postgresdb"
)

// Dependencies holds the dependencies required by the application.
type Dependencies struct {
	CoreValueService    corevalues.Service
	AppreciationService appreciation.Service
	UserService         user.Service
	RewardService       reward.Service
	ReportAppreciationService reportappreciations.Service
}

// NewService initializes and returns a Dependencies instance with the given database connection.
func NewService(db *sqlx.DB) Dependencies {
	// Initialize repository dependencies using the provided database connection.

	coreValueRepo := repository.NewCoreValueRepo(db)
	userRepo := repository.NewUserRepo(db)
	appreciationRepo := repository.NewAppreciationRepo(db)
	rewardRepo := repository.NewRewardRepo(db)
	reportAppreciationRepo := repository.NewReportRepo(db)

	coreValueService := corevalues.NewService(coreValueRepo)
	userService := user.NewService(userRepo)
	appreciationService := appreciation.NewService(appreciationRepo, coreValueRepo)
	rewardService := reward.NewService(rewardRepo,appreciationRepo)
	reportAppreciationService := reportappreciations.NewService(reportAppreciationRepo)

	return Dependencies{
		CoreValueService:    coreValueService,
		AppreciationService: appreciationService,
		UserService:         userService,
		RewardService:       rewardService,
		ReportAppreciationService: reportAppreciationService,
	}

}
