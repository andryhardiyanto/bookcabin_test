package main

import (
	"backend/config"
	"backend/internal/app"
	"backend/internal/respository"
	"backend/internal/service"
	"backend/pkg/db"
	"backend/pkg/fibers"
	"backend/pkg/logger"
	validators "backend/pkg/validator"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	err = logger.InitLogger()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	db, err := db.NewSqlLite()
	if err != nil {
		logger.Fatal("Failed to initialize database", zap.Error(err))
	}

	validator, err := validators.NewValidator()
	if err != nil {
		logger.Fatal("Failed to initialize validator", zap.Error(err))
	}

	repo := respository.NewVoucherRepository(db)
	service := service.NewVoucherService(repo)

	fiberApp := fiber.New(fiber.Config{
		ErrorHandler:  fibers.MiddlewareErrorHandler,
		AppName:       config.Cfg.AppName,
		CaseSensitive: true,
		StrictRouting: false,
	})

	app := app.NewApp(service, fiberApp, validator)

	app.Start()
}
