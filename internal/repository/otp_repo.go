package repository

import (
	"context"
	"errors"
	"github.com/mojtabamovahedi/otp/internal/repository/types"
	"github.com/mojtabamovahedi/otp/pkg/otp"
	"github.com/mojtabamovahedi/otp/pkg/redis"
	"time"
)

var (
	ErrOtpNotFound  = errors.New("OTP not found")
	ErrInvalidOtp   = errors.New("invalid OTP")
	ErrOtpGenerated = errors.New("OTP generated")
)

type OtpRepo interface {
	Generate(ctx context.Context, phoneNumber string) error
	Verify(ctx context.Context, phoneNumber string, otp string) (bool, error)
}

type otpRepo struct {
	oc *redis.ObjectCacher[types.OTP]
}

func NewOtpRepo(oc *redis.ObjectCacher[types.OTP]) OtpRepo {
	return &otpRepo{
		oc: oc,
	}
}

func (o *otpRepo) Generate(ctx context.Context, phoneNumber string) error {
	var err error
	// check phone
	data, err := o.oc.Get(ctx, phoneNumber)
	if len(data.Code) == otp.CodeLength {
		since := time.Since(time.UnixMilli(data.CreatedAt))
		if since <= 2*time.Minute {
			return ErrOtpGenerated
		} else {
			_ = o.oc.Del(ctx, phoneNumber)
		}
	}

	// generate
	code, err := otp.EncodeToString()
	if err != nil {
		return err
	}

	err = o.oc.Set(ctx, phoneNumber, types.NewOTP(code))
	if err != nil {
		return err
	}
	return nil
}

func (o *otpRepo) Verify(ctx context.Context, phoneNumber string, otp string) (bool, error) {
	data, err := o.oc.Get(ctx, phoneNumber)
	if err != nil {
		if errors.Is(err, redis.ErrRedisNotFound) {
			return false, ErrOtpNotFound
		}
		return false, err
	}

	if data.Code != otp {
		return false, ErrInvalidOtp
	}

	_ = o.oc.Del(ctx, phoneNumber)

	return true, nil

}
