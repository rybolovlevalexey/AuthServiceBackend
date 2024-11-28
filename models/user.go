package models

import (
    "gorm.io/gorm"
)

type User struct {
    ID       uint   `gorm:"primaryKey"`
    Username string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
}

// Создание пользователя
func CreateUser(db *gorm.DB, user *User) error {
    return db.Create(user).Error
}

// Получение пользователя по имени
func GetUserByUsername(db *gorm.DB, username string) (*User, error) {
    var user User
    if err := db.Where("username = ?", username).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}