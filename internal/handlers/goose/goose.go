package goose

import (
	"context"
	"embed"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"transactionroutine/internal/database"
)

const (
	PostgresDialect = "postgres"
)

type GooseOption func(conn *database.Database) ([]goose.OptionsFunc, error)

func WithTableName(table string) GooseOption {
	return func(conn *database.Database) ([]goose.OptionsFunc, error) {
		goose.SetTableName(table)
		return nil, nil
	}
}

func WithEmbeddedMigrations(fs embed.FS) GooseOption {
	return func(conn *database.Database) ([]goose.OptionsFunc, error) {
		goose.SetBaseFS(fs)
		return nil, nil
	}
}

func WithAllowMissing() GooseOption {
	return func(conn *database.Database) ([]goose.OptionsFunc, error) {
		return []goose.OptionsFunc{goose.WithAllowMissing()}, nil
	}
}

func EnsureLatest(conn *database.Database, migrationPath string, options ...GooseOption) error {
	gooseOptions := make([]goose.OptionsFunc, 0)
	for _, option := range options {
		gopts, err := option(conn)
		if err != nil {
			return err
		}
		gooseOptions = append(gooseOptions, gopts...)
	}
	err := goose.SetDialect(PostgresDialect)
	if err != nil {
		return err
	}

	err = goose.Up(conn.Conn, migrationPath, gooseOptions...)
	if err != nil {
		return err
	}

	return nil
}

func FullDown(conn *database.Database, migrationPath string, options ...GooseOption) error {
	gooseOptions := make([]goose.OptionsFunc, 0)
	for _, option := range options {
		gopts, err := option(conn)
		if err != nil {
			return err
		}
		gooseOptions = append(gooseOptions, gopts...)
	}
	err := goose.SetDialect(PostgresDialect)
	if err != nil {
		return err
	}
	err = goose.Reset(conn.Conn, migrationPath, gooseOptions...)
	if err != nil {
		return err
	}

	return nil
}

func Migrate(ctx context.Context, db *database.Database, migrationsPath string) error {
	if err := EnsureLatest(db, migrationsPath); err != nil {
		return err
	}
	return nil
}
