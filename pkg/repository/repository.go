package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/novriantama/question1/pkg/models"
	"github.com/novriantama/question1/pkg/sqlc/db"
)

type Repository interface {
	GetUserByID(ctx context.Context, id int) (*db.User, error)
	CreateUser(ctx context.Context, payload models.UserPayload) error
	GenerateOtp(ctx context.Context, payload models.SetOtpPayload) error
	GetUserByPhone(ctx context.Context, phone string) (*db.User, error)
}

type repository struct {
	conn    *pgx.Conn
	queries *db.Queries
}

func NewRepository(conn *pgx.Conn, queries *db.Queries) Repository {
	return &repository{conn: conn, queries: queries}
}

func (r *repository) GetUserByID(ctx context.Context, id int) (*db.User, error) {
	var user db.User
	user, err := r.queries.GetUserByID(ctx, int64(id))
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) CreateUser(ctx context.Context, payload models.UserPayload) error {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	qtx := r.queries.WithTx(tx)
	phoneNumberText := pgtype.Text{payload.PhoneNumber, true}
	nameText := pgtype.Text{payload.Name, true}

	if err := qtx.CreateUser(ctx, db.CreateUserParams{
		Name:        nameText,
		PhoneNumber: phoneNumberText,
	}); err != nil {
		return err
	}
	tx.Commit(ctx)
	return nil
}

func (r *repository) GenerateOtp(ctx context.Context, payload models.SetOtpPayload) error {
	tx, err := r.conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	qtx := r.queries.WithTx(tx)
	phoneNumberText := pgtype.Text{payload.PhoneNumber, true}
	otpText := pgtype.Text{payload.Otp, true}
	otpExpiryTime := pgtype.Timestamp{payload.OtpExpiryTime, 0, true}
	User, err := qtx.GetUserByPhone(ctx, phoneNumberText)
	if err != nil {
		return err
	}
	if err := qtx.UpdateOtp(ctx, db.UpdateOtpParams{
		ID:            User.ID,
		Otp:           otpText,
		OtpExpiryTime: otpExpiryTime,
	}); err != nil {
		return err
	}
	tx.Commit(ctx)
	return nil
}

func (r *repository) GetUserByPhone(ctx context.Context, phone string) (*db.User, error) {
	var user db.User
	phoneNumberText := pgtype.Text{phone, true}
	user, err := r.queries.GetUserByPhone(ctx, phoneNumberText)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
