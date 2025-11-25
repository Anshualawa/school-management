package database

import (
	"fmt"
	"log"

	"github.com/Anshualawa/school-management/internal/config"
	"github.com/Anshualawa/school-management/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect database %w", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Student{},
		&models.Teacher{},
		&models.Class{},
		&models.Subject{},
		&models.ClassSubject{},
		&models.Attendance{},
		&models.Exam{},
		&models.Marks{},
		&models.Book{},
		&models.Fee{},
		&models.FeePayment{},
	)

	if err != nil {
		log.Fatalf("some error in auto migration %s", err.Error())
	}

	log.Println("database connect successfull")
	return db, nil
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("failed to close DB:%v", err)
		return
	}
	sqlDB.Close()
	log.Println("database connection closed")
}
