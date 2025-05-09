package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojtabamovahedi/otp/config"
	"github.com/mojtabamovahedi/otp/internal/service"
)

type Server struct {
	router *gin.Engine
	otpSrv *service.OTPService
	cfg    config.ServerConfig
}

func NewServer(otpSrv *service.OTPService, cfg config.ServerConfig) *Server {
	r := gin.New()
	return &Server{
		router: r,
		otpSrv: otpSrv,
		cfg:    cfg,
	}
}

func (s *Server) Run() error {
	s.router.Use(Logger(), Recovery(), Limiter())

	s.registerRoutes()

	return s.router.Run(fmt.Sprintf(":%d", s.cfg.HttpPort))
}

func (s *Server) registerRoutes() {
	group := s.router.Group("/api/v1/otp")
	otp := newOTP(s.otpSrv)
	group.POST("/generate", otp.GenerateOTP())
	group.POST("/verify", otp.VerifyOTP())
}

func (s *Server) Stop() error {
	return nil
}
