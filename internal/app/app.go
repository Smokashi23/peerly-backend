package app

import (
	"github.com/jmoiron/sqlx"
	"github.com/joshsoftware/peerly-backend/internal/app/appreciation"
	corevalues "github.com/joshsoftware/peerly-backend/internal/app/coreValues"
	reportappreciations "github.com/joshsoftware/peerly-backend/internal/app/reportAppreciations"

	reward "github.com/joshsoftware/peerly-backend/internal/app/reward"
	user "github.com/joshsoftware/peerly-backend/internal/app/users"
	repository "github.com/joshsoftware/peerly-backend/internal/repository/postgresdb"
)

// Dependencies holds the dependencies required by the application.
type Dependencies struct {
	CoreValueService          corevalues.Service
	AppreciationService       appreciation.Service
	UserService               user.Service
	ReportAppreciationService reportappreciations.Service
	RewardService             reward.Service
}

// NewService initializes and returns a Dependencies instance with the given database connection.
func NewService(db *sqlx.DB) Dependencies {
	// Initialize repository dependencies using the provided database connection.

	coreValueRepo := repository.NewCoreValueRepo(db)
	userRepo := repository.NewUserRepo(db)
	reportAppreciationRepo := repository.NewReportRepo(db)
	appreciationRepo := repository.NewAppreciationRepo(db)
	rewardRepo := repository.NewRewardRepo(db)

	coreValueService := corevalues.NewService(coreValueRepo)
	appreciationService := appreciation.NewService(appreciationRepo, coreValueRepo)
	userService := user.NewService(userRepo)
	reportAppreciationService := reportappreciations.NewService(reportAppreciationRepo)
	rewardService := reward.NewService(rewardRepo, appreciationRepo)

	return Dependencies{
		CoreValueService:          coreValueService,
		AppreciationService:       appreciationService,
		UserService:               userService,
		ReportAppreciationService: reportAppreciationService,
		RewardService:             rewardService,
	}

}
