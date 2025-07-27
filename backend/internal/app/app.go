package app

import (
	"backend/config"
	"backend/internal/service"
	httpTransport "backend/internal/transport/http"
	"backend/pkg/fibers"
	"backend/pkg/logger"
	validators "backend/pkg/validator"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"go.uber.org/zap"
)

type App struct {
	voucherHandler *httpTransport.VoucherHandler
	fiberApp       *fiber.App
}

func NewApp(voucherService service.VoucherService, fiberApp *fiber.App, validator validators.Validator) *App {
	return &App{
		voucherHandler: httpTransport.NewVoucherHandler(voucherService, validator),
		fiberApp:       fiberApp,
	}
}

func (a *App) Start() {
	a.fiberApp.Use(fibers.MiddlewareRecovery)
	a.fiberApp.Use(fibers.MiddlewareLogging())

	a.fiberApp.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:  []string{"Accept", "Origin", "Authorization", "Content-Type", "X-Api-Key", "X-CSRF-Token", "X-Access-Token", "X-Request-ID", "X-User-Timezone", "X-Client-Timezone"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        int(12 * time.Hour / time.Second),
	}))

	a.fiberApp.Get("/health", func(c fiber.Ctx) error {
		return fibers.SendResponse(c, nil)
	})
	InitRoute(a)

	a.fiberApp.Use(func(c fiber.Ctx) error {
		return fibers.SendResponse(c, fibers.ErrorPageNotFound)
	}) // => handle page not found

	port := config.Cfg.AppPort
	addr := fmt.Sprintf(":%s", port)
	logger.Info("Starting HTTP server", 
		zap.String("port", port),
		zap.String("address", addr))

	go func() {
		if err := a.fiberApp.Listen(addr, fiber.ListenConfig{
			DisableStartupMessage: true,
		}); err != nil && err != http.ErrServerClosed {
			logger.Error("Listen error", zap.Error(err))
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutdown Server ...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.fiberApp.ShutdownWithContext(shutdownCtx); err != nil {
		logger.Error("Shutdown error", zap.Error(err))
	} else {
		logger.Info("Server shutdown completed successfully")
	}
}

func InitRoute(app *App) {
	api := app.fiberApp.Group("/api")
	{
		api.Post("/check", app.voucherHandler.CheckVoucher)
		api.Post("/generate", app.voucherHandler.GenerateVoucher)
	}
}
