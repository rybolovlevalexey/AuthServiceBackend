package utils

import (
    "github.com/dgrijalva/jwt-go"
    "time"
    "AuthServiceBackend/config"
    "golang.org/x/crypto/bcrypt"
)

// GenerateToken создает новый JWT токен для указанного пользователя.
func GenerateToken(username string) (string, error) {
    claims := jwt.MapClaims{
        "username": username,
        "exp":      time.Now().Add(time.Hour * 1).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(config.JWTSecret))
}

// ValidateToken проверяет действительность JWT токена и возвращает его утверждения.
func ValidateToken(tokenString string) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Проверяем, что метод подписи токена соответствует ожидаемому
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
        }
        return []byte(config.JWTSecret), nil
    })

    if err != nil {
        return nil, err
    }

    // Проверяем, что токен действителен и возвращаем утверждения
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    }

    return nil, jwt.NewValidationError("invalid token", 512)
}


// CheckPasswordHash проверяет, соответствует ли введенный пароль хешу пароля.
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func CreatePasswordHash(password string) (string) {
    bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes)
}
