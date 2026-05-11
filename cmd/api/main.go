package main

import (
	"log"

	"toki/internal/delivery/http"
	"toki/internal/delivery/http/handler"
	"toki/internal/infrastructure/config"
	"toki/internal/infrastructure/database"
	"toki/internal/repository/postgres"
	itemUC "toki/internal/usecase/item"
	stockUC "toki/internal/usecase/stock"
)

func main() {
	cfg := config.Load()

	db := database.NewPostgres(cfg.DBUrl)
	_ = db

	itemRepo := postgres.NewItemRepository(db)
	itemUC := itemUC.NewUsecase(itemRepo, db)
	itemHandler := handler.NewItemHandler(itemUC)

	stockRepo := postgres.NewStockRepository(db)
	stockUC := stockUC.NewUsecase(stockRepo, db)
	stockHandler := handler.NewStockHandler(stockUC)

	app := http.NewRouter(itemHandler, stockHandler)

	log.Println("🚀 Server running on port", cfg.AppPort)
	log.Fatal(app.Listen(":" + cfg.AppPort))
}
