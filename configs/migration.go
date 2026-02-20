package configs

import (
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func RunMigrations() {
	db := DB

	// 1. pastikan tabel schema_migrations ada
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version VARCHAR(255) PRIMARY KEY,
			applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatal("failed to create schema_migrations:", err)
	}

	// 2. ambil semua file .sql dalam folder
	migrationDir := "migrations"

	files, err := os.ReadDir(migrationDir)
	if err != nil {
		log.Fatal("failed to read migration directory:", err)
	}

	// 3. sortir berdasarkan nama file (001, 002, dst)
	var migrations []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") {
			migrations = append(migrations, file.Name())
		}
	}
	sort.Strings(migrations)

	// 4. jalankan satu per satu
	for _, filename := range migrations {
		version := filename // version = nama file

		var exists bool
		err := db.QueryRow(`
			SELECT EXISTS (
				SELECT 1 FROM schema_migrations WHERE version = $1
			)
		`, version).Scan(&exists)

		if err != nil {
			log.Fatal("failed to check migration:", err)
		}

		if exists {
			log.Println("Skipping migration:", version)
			continue
		}

		// baca file SQL
		path := filepath.Join(migrationDir, filename)
		sqlBytes, err := os.ReadFile(path)
		if err != nil {
			log.Fatal("failed to read migration file:", err)
		}

		log.Println("Running migration:", version)

		// eksekusi SQL
		_, err = db.Exec(string(sqlBytes))
		if err != nil {
			log.Fatalf("migration failed (%s): %v", version, err)
		}

		// simpan version
		_, err = db.Exec(`
			INSERT INTO schema_migrations (version)
			VALUES ($1)
		`, version)
		if err != nil {
			log.Fatal("failed to save migration version:", err)
		}

		log.Println("Migration applied:", version)
	}

	log.Println("All migrations completed")
}
