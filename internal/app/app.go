package app

import (
	"github.com/jmoiron/sqlx"
	corevalues "github.com/joshsoftware/peerly-backend/internal/app/coreValues"

	user "github.com/joshsoftware/peerly-backend/internal/app/users"

	repository "github.com/joshsoftware/peerly-backend/internal/repository/postgresdb"
)

// Dependencies holds the dependencies required by the application.
type Dependencies struct {
	CoreValueService corevalues.Service
	UserService      user.Service
}

// NewService initializes and returns a Dependencies instance with the given database connection.
func NewService(db *sqlx.DB) Dependencies {
	// Initialize repository dependencies using the provided database connection.
	coreValueRepo := repository.NewCoreValueRepo(db)
	userRepo := repository.NewUserRepo(db)
	coreValueService := corevalues.NewService(coreValueRepo)
	userService := user.NewService(userRepo)
	return Dependencies{
		CoreValueService: coreValueService,
		UserService:      userService,
	}
}
