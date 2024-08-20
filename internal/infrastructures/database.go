package infrastructures

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"sagara-msib-test/internal/entities"
)

func NewDatabase(ctx context.Context) (db *gorm.DB, err error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database: %w", err)
	}

	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	if err := ensureTableExists(ctx, db); err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to PostgreSQL using GORM!")
	return db, nil
}

func ensureTableExists(ctx context.Context, db *gorm.DB) error {
	// Pastikan tabel `baju` ada
	err := db.WithContext(ctx).AutoMigrate(&entities.Baju{})
	if err != nil {
		return fmt.Errorf("failed to migrate table: %w", err)
	}
	return nil
}
