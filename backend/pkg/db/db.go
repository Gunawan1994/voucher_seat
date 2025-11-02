package db

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
	"voucher_seat/app/domain"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbPg   *gorm.DB
	oncePg sync.Once
)

func NewConn() *gorm.DB {
	var err error

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("POSTGRES_ADDR"),
		"postgres",
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_PORT"),
	)

	oncePg.Do(func() {
		dbPg, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Fatalf("Failed to connect to PostgreSQL: %v", err)
		}

		sqlDB, err := dbPg.DB()
		if err != nil {
			log.Fatalf("Failed to get sql.DB: %v", err)
		}

		// Configure connection pool
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxIdleTime(30 * time.Minute)
		sqlDB.SetConnMaxLifetime(time.Hour)

		// Auto-migrate ensures the table exists
		if err := dbPg.AutoMigrate(&domain.Assignments{}); err != nil {
			log.Fatalf("❌ Failed to migrate assignments table: %v", err)
		}

		log.Println("PostgreSQL connected and 'assignments' table ensured.")
	})

	return dbPg
}
