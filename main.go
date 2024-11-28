package main

import (
    "log"
    "net/http"
    "AuthServiceBackend/config"
    "AuthServiceBackend/routes"
)

func main() {
	// Установка ключей и загрузка конфигурации
	config.GetReadyToWork()
    config.LoadConfig()

    r := routes.SetupRoutes()
    log.Println("Сервер запущен на :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal(err)
    }
}
