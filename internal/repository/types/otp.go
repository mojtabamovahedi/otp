package types

import "time"



var (
	// OTPExpireTime is the expiration time for OTP in seconds
	OTPExpireTime = 2 * 60 // 2 minutes
	
)

type OTP struct {
	Code	 string `json:"code"`
	ExpireAt int64  `json:"expire_at"`
	CreatedAt int64  `json:"created_at"`
}


func NewOTP(code string) OTP {
	now := time.Now()
	return OTP{
		Code:     code,
		ExpireAt: now.Add(time.Duration(OTPExpireTime) * time.Second).Unix(),
		CreatedAt: now.Unix(),
	}
}