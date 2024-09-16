package postgres

import (
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	postgres2 "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDatabaseConnectionString() string {
	return "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=" + os.Getenv("DB_SSLMODE")
}

func NewConnection(shouldMigrate bool) *gorm.DB {
	gormDB, err := gorm.Open(postgres2.Open(getDatabaseConnectionString()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if shouldMigrate {
		gormDB.AutoMigrate(Product{}, Seller{})
	}

	return gormDB
}
