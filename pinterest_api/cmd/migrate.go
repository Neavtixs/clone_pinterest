package main

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
)

func main() {
	// Load konfigurasi dari file config
	v := viper.New()
	v.SetConfigFile("config.json") // Ubah sesuai dengan lokasi file konfigurasi kamu
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	// Ambil koneksi string dari config
	username := v.GetString("database.username")
	password := v.GetString("database.password")
	host := v.GetString("database.host")
	port := v.GetInt("database.port")
	database := v.GetString("database.name")

	var dsn string
	if password == "" {
		dsn = fmt.Sprintf("postgres://%s@%s:%d/%s?sslmode=disable",
			username, host, port, database)
	} else {
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			username, password, host, port, database)
	}

	// Inisialisasi migrasi
	m, err := migrate.New(
		"file://internal/db/migrations", // folder tempat file migrasi .sql disimpan
		dsn,
	)
	if err != nil {
		log.Fatalf("migration init error: %v", err)
	}

	// Jalankan migrasi naik (UP)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration up error: %v", err)
	}

	fmt.Println("✅ Migration completed successfully")
}
