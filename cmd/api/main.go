package main

import (
	"log"

	"toki/internal/delivery/http"
	"toki/internal/delivery/http/handler"
	"toki/internal/infrastructure/config"
	"toki/internal/infrastructure/database"
	"toki/internal/repository/postgres"
	inboundUC "toki/internal/usecase/inbound"
	itemUC "toki/internal/usecase/item"
	reportUC "toki/internal/usecase/report"
	salesUC "toki/internal/usecase/sales"
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

	inboundRepo := postgres.NewInboundRepository()
	inboundUC := inboundUC.NewUsecase(inboundRepo, db)
	inboundHandler := handler.NewInboundHandler(inboundUC)

	salesRepo := postgres.NewSalesRepository()
	salesUC := salesUC.NewUsecase(salesRepo, db)
	salesHandler := handler.NewSalesHandler(salesUC)

	reportRepo := postgres.NewReportRepository(db)
	reportUC := reportUC.NewUsecase(reportRepo)
	reportHandler := handler.NewReportHandler(reportUC)

	app := http.NewRouter(itemHandler, stockHandler, inboundHandler, salesHandler, reportHandler)

	log.Println("🚀 Server running on port", cfg.AppPort)
	log.Fatal(app.Listen(":" + cfg.AppPort))
}
