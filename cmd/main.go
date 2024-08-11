package main

import (
	"log"
	"net/http"
	"strconv"

	"forum/pkg/config"
	"forum/pkg/database"
	"forum/pkg/handlers"
	"forum/pkg/repositories"
	"forum/pkg/services"

	"github.com/gorilla/mux"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	db, err := database.NewSQLiteDB(cfg.DatabaseDSN)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	r := mux.NewRouter()
	r.HandleFunc("/users/register", userHandler.RegisterUser).Methods(http.MethodPost)
	r.HandleFunc("/users/{id:[0-9]+}", userHandler.GetUserByID).Methods(http.MethodGet)

	log.Printf("Server is running on port %d", cfg.ServerPort)
	if err := http.ListenAndServe(":"+strconv.Itoa(cfg.ServerPort), r); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
