package db

import (
	"backend/config"
	"backend/pkg/logger"
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
)

type sqlLite struct {
	DB *sqlx.DB
}

type DB interface {
	Select(ctx context.Context, query string, dest any, args ...any) error
	Get(ctx context.Context, query string, dest any, args ...any) error
	Insert(ctx context.Context, query string, args ...any) error
	Delete(ctx context.Context, query string, args ...any) error
	Update(ctx context.Context, query string, args ...any) error
}

func NewSqlLite() (DB, error) {
	db, err := sqlx.Open("sqlite3", config.Cfg.DBPath)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(config.Cfg.DBMaxOpenConns)
	db.SetMaxIdleConns(config.Cfg.DBMaxIdleConns)
	db.SetConnMaxIdleTime(config.Cfg.DBConnMaxIdleTime)
	db.SetConnMaxLifetime(config.Cfg.DBConnMaxLifetime)

	initMigration(db)

	logger.Info("Database initialized successfully", 
		zap.String("db_path", config.Cfg.DBPath))

	return &sqlLite{DB: db}, nil
}

func initMigration(db *sqlx.DB) {
	createTable := `
    CREATE TABLE IF NOT EXISTS vouchers (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        crew_name TEXT,
        crew_id TEXT,
        flight_number TEXT,
        flight_date TEXT,
        aircraft_type TEXT,
        seat1 TEXT,
        seat2 TEXT,
        seat3 TEXT,
        created_at TEXT
    );
    `
	_, err := db.Exec(createTable)
	if err != nil {
		logger.Fatal("Failed to create table", zap.Error(err))
	}

	// add unique index for flight_number & flight_date
	createIndex := `
	CREATE UNIQUE INDEX IF NOT EXISTS idx_flight_number_date ON vouchers (flight_number, flight_date);
	`

	_, err = db.Exec(createIndex)
	if err != nil {
		logger.Fatal("Failed to create index", zap.Error(err))
	}

	logger.Debug("Database migration completed successfully")
}

func (s *sqlLite) Select(ctx context.Context, query string, dest any, args ...any) error {
	return s.DB.SelectContext(ctx, dest, query, args...)
}

func (s *sqlLite) Get(ctx context.Context, query string, dest any, args ...any) error {
	return s.DB.GetContext(ctx, dest, query, args...)
}

func (s *sqlLite) Insert(ctx context.Context, query string, args ...any) error {
	_, err := s.DB.ExecContext(ctx, query, args...)
	return err
}

func (s *sqlLite) Delete(ctx context.Context, query string, args ...any) error {
	_, err := s.DB.ExecContext(ctx, query, args...)
	return err
}

func (s *sqlLite) Update(ctx context.Context, query string, args ...any) error {
	_, err := s.DB.ExecContext(ctx, query, args...)
	return err
}

func (s *sqlLite) Close() error {
	return s.DB.Close()
}
