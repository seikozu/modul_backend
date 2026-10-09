package config

import (
	"errors"
	"log/slog"

	"github.com/gofiber/fiber/v2"

	"modul7/app/model"
	"modul7/helper"
	"modul7/middleware"
	"modul7/route"
)

func NewApp(logger *slog.Logger, deps route.Dependencies) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      GetEnv("APP_NAME", "Praktikum Backend Lanjut"),
		ErrorHandler: newErrorHandler(logger),
		BodyLimit:    1 * 1024 * 1024, // 1 MB
	})

	middleware.Register(app, logger, GetEnv("ALLOWED_ORIGINS", ""))
	route.Register(app, deps)

	app.Use(func(c *fiber.Ctx) error {
		return helper.Fail(c, fiber.StatusNotFound, "endpoint tidak ditemukan")
	})

	return app
}

func newErrorHandler(logger *slog.Logger) fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		requestID := helper.RequestID(c)
		status := fiber.StatusInternalServerError
		message := "terjadi error pada server"
		code := "INTERNAL_ERROR"
		fields := map[string]string(nil)

		var appErr *helper.AppError
		switch {
		case errors.As(err, &appErr):
			status = appErr.Status
			code = appErr.Code
			message = appErr.Message
			fields = appErr.Fields
		case errors.Is(err, fiber.ErrRequestEntityTooLarge):
			status = fiber.StatusRequestEntityTooLarge
			code = "PAYLOAD_TOO_LARGE"
			message = "ukuran body melebihi batas yang diizinkan"
		default:
			var fiberErr *fiber.Error
			if errors.As(err, &fiberErr) {
				status = fiberErr.Code
				code = "HTTP_ERROR"
				message = fiberErr.Message
			}
		}

		if status >= fiber.StatusInternalServerError {
			logger.Error("request_failed",
				slog.String("request_id", requestID),
				slog.String("path", c.Path()),
				slog.String("code", code),
				slog.Int("status", status),
				slog.String("error", err.Error()))
		} else {
			logger.Warn("request_rejected",
				slog.String("request_id", requestID),
				slog.String("path", c.Path()),
				slog.String("code", code),
				slog.Int("status", status))
		}

		return c.Status(status).JSON(model.ErrorResponse{
			Success:   false,
			Code:      code,
			Message:   message,
			Fields:    fields,
			RequestID: requestID,
		})
	}
}
