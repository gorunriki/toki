package main

import (
	"log"

	"toki/internal/delivery/http"
	"toki/internal/delivery/http/handler"
	"toki/internal/infrastructure/config"
	"toki/internal/infrastructure/database"
	"toki/internal/repository/postgres"

	authUC "toki/internal/usecase/auth"
	inboundUC "toki/internal/usecase/inbound"
	itemUC "toki/internal/usecase/item"
	reportUC "toki/internal/usecase/report"
	salesUC "toki/internal/usecase/sales"
	stockUC "toki/internal/usecase/stock"
)

func main() {
	cfg := config.Load()

	db := database.NewPostgres(cfg.DBUrl)

	// ITEM
	itemRepo := postgres.NewItemRepository(db)

	itemUsecase := itemUC.NewUsecase(
		itemRepo,
		db,
	)

	itemHandler := handler.NewItemHandler(
		itemUsecase,
	)

	// STOCK
	stockRepo := postgres.NewStockRepository(
		db,
	)

	stockUsecase := stockUC.NewUsecase(
		stockRepo,
	)

	stockHandler := handler.NewStockHandler(
		stockUsecase,
	)

	// INBOUND
	inboundRepo := postgres.NewInboundRepository()

	inboundUsecase := inboundUC.NewUsecase(
		inboundRepo,
		stockRepo,
		db,
	)

	inboundHandler := handler.NewInboundHandler(
		inboundUsecase,
	)

	// SALES
	salesRepo := postgres.NewSalesRepository()

	salesUsecase := salesUC.NewUsecase(
		salesRepo,
		db,
	)

	salesHandler := handler.NewSalesHandler(
		salesUsecase,
	)

	// REPORT
	reportRepo := postgres.NewReportRepository(
		db,
	)

	reportUsecase := reportUC.NewUsecase(
		reportRepo,
	)

	reportHandler := handler.NewReportHandler(
		reportUsecase,
	)

	// AUTH
	userRepo := postgres.NewUserRepository(
		db,
	)

	authUsecase := authUC.NewUsecase(
		userRepo,
	)

	authHandler := handler.NewAuthHandler(
		authUsecase,
	)

	// ROUTER
	app := http.NewRouter(
		itemHandler,
		stockHandler,
		inboundHandler,
		salesHandler,
		reportHandler,
		authHandler,
	)

	log.Println(
		"🚀 Server running on port",
		cfg.AppPort,
	)

	log.Fatal(
		app.Listen(":" + cfg.AppPort),
	)
}
