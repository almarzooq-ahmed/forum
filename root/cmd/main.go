package main

import (
	"log"
	"strconv"

	config "forum/root/internal/config"
	handlers "forum/root/internal/delivery/http/handlers"
	router "forum/root/internal/delivery/http/routers"
	repositories "forum/root/internal/domain/repositories"
	usecases "forum/root/internal/domain/usecases"
	database "forum/root/internal/infrastructure/database"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	db, err := database.CreateDB(cfg.DatabaseDSN)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	userRepo := repositories.NewUserRepository(db)
	userUseCase := usecases.NewUserUseCase(userRepo)
	userHandler := handlers.NewUserHandler(userUseCase)

	router := router.SetupRouter(userHandler)

	err = router.Run(":" + strconv.Itoa(cfg.ServerPort))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
