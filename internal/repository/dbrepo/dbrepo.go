package dbrepo

import (
	"github.com/y-kouhei9/bookings-app/internal/repository"
	"database/sql"
	"github.com/y-kouhei9/bookings-app/internal/config"
	
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

// NewPostgresRepo returns postgredDBRepo struct
func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB: conn,
	}
}
