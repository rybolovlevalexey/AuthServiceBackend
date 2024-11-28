package database

import (
	"AuthServiceBackend/config"
	"AuthServiceBackend/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    var err error
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", 
        config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Ошибка подключения к базе данных:", err)
    }

    // Автоматическая миграция
    err = DB.AutoMigrate(&models.User{})
    if err != nil {
        log.Fatal("Ошибка миграции:", err)
    }

    log.Println("База данных успешно подключена и миграция выполнена")
}
