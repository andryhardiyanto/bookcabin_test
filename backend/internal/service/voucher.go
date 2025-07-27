package service

import (
	"backend/internal/model"
	"backend/internal/respository"
	"backend/pkg/fibers"
	"context"
)

type voucherService struct {
	repo respository.VoucherRepository
}

type VoucherService interface {
	CheckVoucher(ctx context.Context, req *model.RequestCheckVoucher) (*model.ResponseCheckVoucher, error)
	GenerateVoucher(ctx context.Context, req *model.RequestGenerateVoucher) (*model.ResponseGenerateVoucher, error)
}

func NewVoucherService(repo respository.VoucherRepository) VoucherService {
	return &voucherService{repo: repo}
}

func (v *voucherService) CheckVoucher(ctx context.Context, req *model.RequestCheckVoucher) (*model.ResponseCheckVoucher, error) {
	err := v.repo.CheckVoucher(ctx, req)
	if err != nil {
		if err == fibers.ErrorVoucherAlreadyExist {
			return &model.ResponseCheckVoucher{Exists: true}, nil
		}
		return nil, err
	}

	return &model.ResponseCheckVoucher{Exists: false}, nil
}

func (v *voucherService) GenerateVoucher(ctx context.Context, req *model.RequestGenerateVoucher) (*model.ResponseGenerateVoucher, error) {
	resp, err := v.CheckVoucher(ctx, &model.RequestCheckVoucher{
		FlightNumber: req.FlightNumber,
		Date:         req.Date,
	})
	if err != nil {
		return nil, err
	}

	if resp.Exists {
		return nil, fibers.ErrorVoucherAlreadyExist
	}

	seat1, seat2, seat3 := model.GenerateRandomSeatsByAircraftType(req.Aircraft)
	err = v.repo.GenerateVoucher(ctx, &model.RequestDatabaseGenerateVoucher{
		CrewName:     req.Name,
		CrewID:       req.ID,
		FlightNumber: req.FlightNumber,
		Date:         req.Date,
		AircraftType: req.Aircraft,
		Seat1:        seat1,
		Seat2:        seat2,
		Seat3:        seat3,
	})
	if err != nil {
		return nil, err
	}

	return &model.ResponseGenerateVoucher{
		Seats: []string{seat1, seat2, seat3},
	}, nil
}
