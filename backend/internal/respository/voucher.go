package respository

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"backend/internal/model"
	"backend/pkg/db"
	"backend/pkg/fibers"
)

type voucherRepository struct {
	db db.DB
}

type VoucherRepository interface {
	CheckVoucher(ctx context.Context, req *model.RequestCheckVoucher) error
	GenerateVoucher(ctx context.Context, req *model.RequestDatabaseGenerateVoucher) error
}

func NewVoucherRepository(db db.DB) VoucherRepository {
	return &voucherRepository{db: db}
}

func (v *voucherRepository) CheckVoucher(ctx context.Context, req *model.RequestCheckVoucher) error {
	query := `SELECT * FROM vouchers WHERE flight_number = ? and flight_date = ? limit 1;`
	voucher := &model.Voucher{}
	err := v.db.Get(ctx, query, voucher, req.FlightNumber, req.Date)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return err
	}

	return fibers.ErrorVoucherAlreadyExist
}

func (v *voucherRepository) GenerateVoucher(ctx context.Context, req *model.RequestDatabaseGenerateVoucher) error {
	query := `INSERT INTO vouchers (crew_name, crew_id, flight_number, flight_date, aircraft_type, seat1, seat2, seat3, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP);`
	err := v.db.Insert(ctx, query, req.CrewName, req.CrewID, req.FlightNumber, req.Date, req.AircraftType, req.Seat1, req.Seat2, req.Seat3)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return fibers.ErrorVoucherAlreadyExist
		}
		return err
	}

	return nil
}
