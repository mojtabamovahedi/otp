package http

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/mojtabamovahedi/otp/internal/service"
	"github.com/mojtabamovahedi/otp/pkg/otp"
	"net/http"
	"regexp"
)

type otpApi struct {
	srv *service.OTPService
}

var phoneRegex = regexp.MustCompile("^[0-9]{10}$")

func newOTP(svc *service.OTPService) *otpApi {
	return &otpApi{
		srv: svc,
	}
}

func (o *otpApi) GenerateOTP() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			err error
			ctx = c.Request.Context()
		)

		req := GenerateOTPReqBody{}

		err = c.ShouldBindBodyWith(&req, binding.JSON)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "wrong body",
			})
			return
		}

		if !phoneRegex.MatchString(req.Phone) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "wrong phone",
			})
			return
		}

		err = o.srv.GenerateOtp(ctx, req.Phone)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "OTP generated",
		})
		return
	}
}

func (o *otpApi) VerifyOTP() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			err error
			ctx = c.Request.Context()
		)

		req := VerifyOTPReqBody{}
		err = c.ShouldBindBodyWith(&req, binding.JSON)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "wrong body",
			})
			return
		}

		if !phoneRegex.MatchString(req.Phone) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "wrong phone",
			})
			return
		}

		if len(req.Code) != otp.CodeLength {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "wrong code",
			})
			return
		}

		_, err = o.srv.VerifyOtp(ctx, req.Phone, req.Code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "OTP verified",
		})
		return
	}
}

type GenerateOTPReqBody struct {
	Phone string `json:"phone" binding:"required"`
}

type VerifyOTPReqBody struct {
	Phone string `json:"phone" binding:"required"`
	Code  string `json:"code" binding:"required"`
}
