package middleware

import (
	"AuthServiceBackend/utils"
	"context"
	"net/http"
	"strings"
)

// AuthMiddleware проверяет наличие и действительность JWT токена в заголовке Authorization.
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Получаем токен из заголовка Authorization
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Отсутствует токен", http.StatusUnauthorized)
            return
        }

        // Удаляем префикс "Bearer " из токена, если он присутствует
        if strings.HasPrefix(token, "Bearer ") {
            token = strings.TrimPrefix(token, "Bearer ")
        } else {
            
        }

        // Проверяем токен
        claims, err := utils.ValidateToken(token)
        if err != nil {
            http.Error(w, "Неверный токен", http.StatusUnauthorized)
            return
        }


        ctx := context.WithValue(r.Context(), "claims", claims)
        r = r.WithContext(ctx)

        // Передаем управление следующему обработчику
        next.ServeHTTP(w, r)
    })
}

