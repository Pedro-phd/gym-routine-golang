package database

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/Pedro-phd/gym-routine-golang/pkg/log"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

type Database struct {
	DB  *sql.DB
	Log *log.CustomLog
}

func InitDatabase(CustomLog *log.CustomLog) *Database {
	godotenv.Load()
	database := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", database)
	if err != nil {
		panic(err)
	}

	db.Ping()
	if err != nil {
		panic(err)
	}

	CustomLog.Success("Connected to database")
	return &Database{DB: db, Log: CustomLog}
}

func (db *Database) RunMigrations() {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		db.Log.Fatal(err.Error())
		panic(err)
	}
	execPath, _ := os.Getwd()
	migrationPath := filepath.Join(execPath, "/internal/database/migrations")

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationPath,
		"postgres", driver)

	if err != nil {
		db.Log.Fatal(err.Error())
		panic(err)
	}

	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			db.Log.Info("No new migrations to apply")
		} else {
			db.Log.Fatal("Failed to apply migrations: " + err.Error())
		}
		return
	}

	db.Log.Success("Migrations applied successfully")
}
