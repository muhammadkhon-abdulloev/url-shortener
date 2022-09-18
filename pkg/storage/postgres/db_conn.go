package postgres

import (
	"fmt"
	"time"

	_ "github.com/jackc/pgx/stdlib" // pgx driver
	"github.com/jmoiron/sqlx"

	"github.com/muhammadkhon-abdulloev/url-shortener/config"
)

const (
	maxOpenConns    = 5
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
	pgDriver        = "pgx"
)

func NewPsqlDB(cfg *config.Config) (*sqlx.DB, error) {
	var dataSource string
	if cfg.Server.Mode == "Development" {
		dataSource = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s",
			cfg.PostgresDev.Host,
			cfg.PostgresDev.Username,
			cfg.PostgresDev.DBName,
			cfg.PostgresDev.SslMode,
			cfg.PostgresDev.Password,
		)
	} else {
		dataSource = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s",
			cfg.Postgres.Host,
			cfg.Postgres.Username,
			cfg.Postgres.DBName,
			cfg.Postgres.SslMode,
			cfg.Postgres.Password,
		)
	}

	db, err := sqlx.Connect(pgDriver, dataSource)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifetime * time.Second)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
