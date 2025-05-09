package service

import (
	"context"

	"github.com/mojtabamovahedi/otp/internal/repository"
)


type OTPService struct {
	repo repository.OtpRepo
}


func NewOTPService(repo repository.OtpRepo) *OTPService {
	return &OTPService{
		repo: repo,
	}
}

func (s *OTPService) GenerateOtp(ctx context.Context, phoneNumber string) error {
	return s.repo.Generate(ctx, phoneNumber)
}

func (s *OTPService) VerifyOtp(ctx context.Context, phoneNumber string, otp string) (bool, error) {
	return s.repo.Verify(ctx, phoneNumber, otp)
}
