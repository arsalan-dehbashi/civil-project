package db

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"

	"gorm.io/gorm"
)

//go:embed migrations/*.up.sql
var migrationFiles embed.FS

func RunMigrations(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB from gorm: %w", err)
	}

	if err := ensureMigrationTable(sqlDB); err != nil {
		return err
	}

	files, err := fs.ReadDir(migrationFiles, "migrations")
	if err != nil {
		return fmt.Errorf("failed to read migration files: %w", err)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".up.sql") {
			continue
		}

		version := strings.TrimSuffix(file.Name(), ".up.sql")
		applied, err := migrationApplied(sqlDB, version)
		if err != nil {
			return err
		}
		if applied {
			continue
		}

		content, err := migrationFiles.ReadFile(filepath.Join("migrations", file.Name()))
		if err != nil {
			return fmt.Errorf("failed to read migration %s: %w", file.Name(), err)
		}

		if err := applyMigration(sqlDB, version, string(content)); err != nil {
			return err
		}
	}

	return nil
}

func ensureMigrationTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version VARCHAR(255) NOT NULL PRIMARY KEY,
			applied_at DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
	`)
	if err != nil {
		return fmt.Errorf("failed to ensure schema_migrations table: %w", err)
	}
	return nil
}

func migrationApplied(db *sql.DB, version string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM schema_migrations WHERE version = ?", version).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to check migration %s: %w", version, err)
	}
	return count > 0, nil
}

func applyMigration(db *sql.DB, version, content string) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start migration %s: %w", version, err)
	}
	defer tx.Rollback()

	for _, statement := range splitSQLStatements(content) {
		if _, err := tx.Exec(statement); err != nil {
			return fmt.Errorf("failed to apply migration %s: %w", version, err)
		}
	}

	if _, err := tx.Exec("INSERT INTO schema_migrations (version) VALUES (?)", version); err != nil {
		return fmt.Errorf("failed to record migration %s: %w", version, err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit migration %s: %w", version, err)
	}

	return nil
}

func splitSQLStatements(content string) []string {
	lines := strings.Split(content, "\n")
	cleaned := make([]string, 0, len(lines))
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "--") {
			continue
		}
		cleaned = append(cleaned, line)
	}

	parts := strings.Split(strings.Join(cleaned, "\n"), ";")
	statements := make([]string, 0, len(parts))
	for _, part := range parts {
		statement := strings.TrimSpace(part)
		if statement != "" {
			statements = append(statements, statement)
		}
	}

	return statements
}
