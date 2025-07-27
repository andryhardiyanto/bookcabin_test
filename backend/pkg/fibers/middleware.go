package fibers

import (
	"backend/pkg/logger"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

func MiddlewareRecovery(c fiber.Ctx) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
			logger.Error("Panic recovered",
				zap.String("error", err.Error()),
				zap.String("method", c.Method()),
				zap.String("path", c.Path()),
				zap.String("ip", c.IP()),
				zap.Any("panic", r))
			
			err = NewError(500, err.Error(), "internal server error")
			_ = SendResponse(c, err)
		}
	}()
	return c.Next()
}

func MiddlewareErrorHandler(c fiber.Ctx, err error) error {
	if err == nil {
		return nil
	}

	defer func() {
		if r := recover(); r != nil {
			// Return default error response
			panicErr := fmt.Errorf("%v", r)
			logger.Error("Panic in error handler",
				zap.String("error", panicErr.Error()),
				zap.String("method", c.Method()),
				zap.String("path", c.Path()),
				zap.String("ip", c.IP()),
				zap.Any("panic", r))
			
			err = NewError(500, panicErr.Error(), "internal server error")
			_ = SendResponse(c, err)
		}
	}()

	if err, ok := err.(*Error); ok {
		return SendResponse(c, err)
	}

	return SendResponse(c, ErrorInternalServerError)
}

func MiddlewareLogging() fiber.Handler {
	return func(c fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		stop := time.Now()
		latency := stop.Sub(start)

		// Base fields untuk semua log
		fields := []zap.Field{
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Duration("latency", latency),
			zap.String("ip", c.IP()),
			zap.String("user_agent", c.Get("User-Agent")),
		}

		if err != nil {
			if errWithStack, ok := err.(*Error); ok {
				fields = append(fields, 
					zap.Int("status_code", errWithStack.Code),
					zap.String("error", errWithStack.Error()))
				
				if errWithStack.Code >= 500 {
					logger.Error("HTTP Request Error", fields...)
				} else {
					logger.Warn("HTTP Request Warning", fields...)
				}
			} else {
				fields = append(fields, 
					zap.Int("status_code", 500),
					zap.String("error", err.Error()))
				logger.Error("HTTP Request Error", fields...)
			}
		} else {
			fields = append(fields, 
				zap.Int("status_code", c.Response().StatusCode()))
			logger.Info("HTTP Request", fields...)
		}

		return err
	}
}
