package database

import (
	"fmt"

	"github.com/khishh/personal-finance-app/graph/model"
	"github.com/khishh/personal-finance-app/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

func NewConnection(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&model.User{})
}
