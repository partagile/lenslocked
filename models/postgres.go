package models

import (
	"database/sql"
	"fmt"
	"io/fs"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
)

// Open will open a SQL connection with the provided
// Postgres database. If you call this, you will need
// to ensure connection is closed with db.Close() method.
func Open(config PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", config.String())

	if err != nil {
		return nil, fmt.Errorf("open %w", err)
	}
	return db, nil
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)
}

func Migrate(db *sql.DB, dir string) error {
	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("migrate dialect: %w", err)
	}
	err = goose.Up(db, dir)
	if err != nil {
		return fmt.Errorf("migrate up: %w", err)
	}
	return nil
}

func MigrateFS(db *sql.DB, migrationsFS fs.FS, dir string) error {
	if dir == "" {
		dir = "."
	}
	goose.SetBaseFS(migrationsFS)
	defer func() {
		goose.SetBaseFS(nil)
	}()
	return Migrate(db, dir)
}
