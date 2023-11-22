package services

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/novriantama/question1/pkg/models"
	"github.com/novriantama/question1/pkg/repository"
	"github.com/novriantama/question1/pkg/sqlc/db"
)

type Service interface {
	GetUserByID(ctx context.Context, id int) (*db.User, error)
	CreateUser(ctx context.Context, payload models.UserPayload) error
	GenerateOtp(ctx context.Context, payload models.SetOtpPayload) error
	VerifyOtp(ctx context.Context, payload models.GetOtpPayload) error
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetUserByID(ctx context.Context, id int) (*db.User, error) {
	return s.repo.GetUserByID(ctx, id)
}

func (s *service) CreateUser(ctx context.Context, payload models.UserPayload) error {
	return s.repo.CreateUser(ctx, payload)
}

func (s *service) GenerateOtp(ctx context.Context, payload models.SetOtpPayload) error {
	rand.Seed(time.Now().UnixNano())
	payload.Otp = strconv.Itoa(rand.Intn(9000) + 1000)
	payload.OtpExpiryTime = time.Now().UTC().Add(time.Minute)
	return s.repo.GenerateOtp(ctx, payload)
}

func (s *service) VerifyOtp(ctx context.Context, payload models.GetOtpPayload) error {
	user, err := s.repo.GetUserByPhone(ctx, payload.PhoneNumber)
	if err != nil {
		return err
	}
	if user.Otp.String != payload.Otp {
		return fmt.Errorf("otp not match")
	}
	if user.OtpExpiryTime.Time.Before(time.Now().UTC()) {
		return fmt.Errorf("otp expired")
	}
	return nil
}
