package db

import (
	"database/sql"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	_ "github.com/lib/pq"
)

func Open(dsn string) *sql.DB {
	log.Println("🔌 Connecting to database...")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Connected to database successfully")
	if err := runMigration(db); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	return db
}

func runMigration(db *sql.DB) error {
	migPath := filepath.Join("db", "sql_migrations", "migrate.sql")
	content, err := ioutil.ReadFile(migPath)
	if err != nil {
		log.Printf(" No migration file found at %s, skipping migration", migPath)
		return nil
	}

	sqlText := string(content)
	upSQL := extractUpSQL(sqlText)
	if upSQL == "" {
		log.Println(" No '-- +migrate Up' section found in migrate.sql")
		return nil
	}

	log.Println("Running migration from migrate.sql ...")
	if _, err := db.Exec(upSQL); err != nil {
		if strings.Contains(err.Error(), "already exists") {
			log.Println("Tables already exist — skipping migration")
			return nil
		}
		return err
	}

	log.Println("Migration executed successfully")
	return nil
}

func extractUpSQL(content string) string {
	parts := strings.Split(content, "-- +migrate Up")
	if len(parts) < 2 {
		return ""
	}
	upSection := strings.Split(parts[1], "-- +migrate Down")
	return strings.TrimSpace(upSection[0])
}
