package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"AuthServiceBackend/database"
	"AuthServiceBackend/middleware"
	"AuthServiceBackend/models"
	"AuthServiceBackend/utils"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
    r := mux.NewRouter()
    database.Connect()

    r.HandleFunc("/login", LoginHandler).Methods("POST")
    r.HandleFunc("/register", RegisterHandler).Methods("POST")
    r.Handle("/protected", middleware.AuthMiddleware(http.HandlerFunc(ProtectedHandler))).Methods("GET")

    return r
}


func LoginHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Получен запрос на авторизацию")

    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Неверный запрос", http.StatusBadRequest)
        return
    }

    // Получаем пользователя из базы данных
    dbUser, err := models.GetUserByUsername(database.DB, user.Username)
    if err != nil || dbUser == nil || !utils.CheckPasswordHash(user.Password, dbUser.Password) {
        http.Error(w, "Неверные учетные данные", http.StatusUnauthorized)
        return
    }

    // Генерируем JWT токен
    token, err := utils.GenerateToken(user.Username)
    if err != nil {
        http.Error(w, "Ошибка генерации токена", http.StatusInternalServerError)
        return
    }

    // Устанавливаем заголовок Content-Type
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"token": token})

    fmt.Println("Авторизация выполнена успешно")
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Получен запрос на доступ к защищённому ресурсу")

    w.Header().Set("Content-Type", "text/plain")
    w.Write([]byte("Это защищенный ресурс"))

    fmt.Println("Доступ к защищённому ресурсу выполнен успешно")
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Получен запрос на регистрацию")

    var user models.User
    //fmt.Print(r.Body)
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Неверный запрос", http.StatusBadRequest)
        return
    }

    if res, _ := models.GetUserByUsername(database.DB, user.Username); res != nil {
        http.Error(w, "Пользователь с таким именем уже существует", http.StatusInternalServerError)
        return
    }

    // Сохраняем пользователя в базе данных
    user.Password = utils.CreatePasswordHash(user.Password)
    if err := models.CreateUser(database.DB, &user); err != nil {
        http.Error(w, "Ошибка создания пользователя", http.StatusInternalServerError)
        return
    }

    fmt.Println("Регистрация выполнена успешно")
}
