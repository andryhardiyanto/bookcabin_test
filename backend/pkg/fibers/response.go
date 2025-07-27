package fibers

import (
	"github.com/gofiber/fiber/v3"
)

type Response struct {
	Success    bool        `json:"success"`
	Data       any         `json:"data,omitempty"`
	Message    string      `json:"message,omitempty"`
	Type       string      `json:"type,omitempty"`
	Violations []Violation `json:"violations,omitempty"`
}

func SendResponse(c fiber.Ctx, data any) error {
	resp := &Response{
		Success: true,
		Data:    data,
	}

	if data != nil {
		if customErr, ok := data.(*Error); ok {
			resp = &Response{
				Success:    false,
				Message:    customErr.Message,
				Type:       customErr.Type,
				Violations: customErr.Violations,
			}
			return c.Status(customErr.Code).JSON(resp)
		}
	}

	return c.Status(int(200)).JSON(resp)
}
