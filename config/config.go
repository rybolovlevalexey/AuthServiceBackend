package config

import (
    "log"
    "os"
)

var (
    JWTSecret   string
    DBHost      string
    DBPort      string
    DBUser      string
    DBPassword  string
    DBName      string
)

func LoadConfig() {
    JWTSecret = os.Getenv("JWT_SECRET")
    if JWTSecret == "" {
        log.Fatal("JWT_SECRET не установлен")
    }

    DBHost = os.Getenv("DB_HOST")
    if DBHost == "" {
        log.Fatal("DB_HOST не установлен")
    }

    DBPort = os.Getenv("DB_PORT")
    if DBPort == "" {
        log.Fatal("DB_PORT не установлен")
    }

    DBUser = os.Getenv("DB_USER")
    if DBUser == "" {
        log.Fatal("DB_USER не установлен")
    }

    DBPassword = os.Getenv("DB_PASSWORD")
    if DBPassword == "" {
        log.Fatal("DB_PASSWORD не установлен")
    }

    DBName = os.Getenv("DB_NAME")
    if DBName == "" {
        log.Fatal("DB_NAME не установлен")
    }
}

func GetReadyToWork() {
    // Устанавливаем значения по умолчанию, если переменные окружения не установлены
    if os.Getenv("JWT_SECRET") == "" {
        os.Setenv("JWT_SECRET", "example_secret") // Здесь должен быть установлен реальный JWT secret
    }

    if os.Getenv("DB_HOST") == "" {
        os.Setenv("DB_HOST", "localhost") // Здесь должен быть установлен реальный хост
    }

    if os.Getenv("DB_PORT") == "" {
        os.Setenv("DB_PORT", "5432") // Здесь должен быть установлен реальный порт
    }

    if os.Getenv("DB_USER") == "" {
        os.Setenv("DB_USER", "postgres") // Здесь должен быть установлен реальный пользователь
    }

    if os.Getenv("DB_PASSWORD") == "" {
        os.Setenv("DB_PASSWORD", "141421") // Здесь должен быть установлен реальный пароль
    }

    if os.Getenv("DB_NAME") == "" {
        os.Setenv("DB_NAME", "auth_db") // Здесь должно быть установлено реальное название базы данных
    }
}
