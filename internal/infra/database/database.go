package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/hamidgh01/Go-Blog-API/config"
	"github.com/hamidgh01/Go-Blog-API/pkg/logging"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(dbConf *config.PostgresConf) (*sql.DB, error) {
	var err error
	logger := logging.GetLogger()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		dbConf.Host,
		dbConf.User,
		dbConf.Password,
		dbConf.DBName,
		dbConf.Port,
	)
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("postgres ping failed: %w", err)
	}

	// sqlDB.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	// sqlDB.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	// sqlDB.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Minute)

	logger.Info("Database connection established successfully.")
	return db, nil
}
