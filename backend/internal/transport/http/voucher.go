package http

import (
	"backend/internal/model"
	"backend/internal/service"
	"backend/pkg/fibers"
	validators "backend/pkg/validator"
	"encoding/json"

	"github.com/gofiber/fiber/v3"
)

type VoucherHandler struct {
	service   service.VoucherService
	validator validators.Validator
}

func NewVoucherHandler(service service.VoucherService, validator validators.Validator) *VoucherHandler {
	return &VoucherHandler{service: service, validator: validator}
}

func (h *VoucherHandler) CheckVoucher(c fiber.Ctx) error {
	req := &model.RequestCheckVoucher{}
	if err := json.Unmarshal(c.Body(), req); err != nil {
		return err
	}

	if err := h.validator.Validate(req); err != nil {
		return err
	}

	resp, err := h.service.CheckVoucher(c.RequestCtx(), req)
	if err != nil {
		return err
	}

	return fibers.SendResponse(c, resp)
}

func (h *VoucherHandler) GenerateVoucher(c fiber.Ctx) error {
	req := &model.RequestGenerateVoucher{}
	if err := json.Unmarshal(c.Body(), req); err != nil {
		return err
	}

	if err := h.validator.Validate(req); err != nil {
		return err
	}

	resp, err := h.service.GenerateVoucher(c.RequestCtx(), req)
	if err != nil {
		return err
	}

	return fibers.SendResponse(c, resp)
}
