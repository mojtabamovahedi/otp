package main

import (
	"context"
	"flag"
	"github.com/mojtabamovahedi/otp/api/handler/http"
	"github.com/mojtabamovahedi/otp/config"
	"github.com/mojtabamovahedi/otp/internal/repository"
	"github.com/mojtabamovahedi/otp/internal/repository/types"
	"github.com/mojtabamovahedi/otp/internal/service"
	"github.com/mojtabamovahedi/otp/pkg/redis"
	"go.uber.org/fx"
)

var cfgPath = flag.String("config", "config.yaml", "Path to config file")

func main() {

	fx.New(
		fx.Invoke(func() {
			flag.Parse()
		}),
		fx.Provide(
			// config.NewConfig,
			provideConfig,

			// redis.NewRedisClient
			provideRedisConn,

			// redis.NewObjectCacher
			provideObjectCacher[types.OTP],

			// repository.NewOtpRepo
			provideOTPRepo,

			// service.NewOTPService
			provideOTPService,
		),
		fx.Invoke(startHttpServer),
	).Run()
}

func startHttpServer(lifecycle fx.Lifecycle, otpSrv *service.OTPService, cfg *config.Config) {
	server := http.NewServer(otpSrv, cfg.Server)
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				return server.Run()
			},
			OnStop: func(ctx context.Context) error {
				return server.Stop()
			},
		},
	)
}

func provideConfig() *config.Config {
	cfg := config.MustReadConfig(*cfgPath)
	return &cfg
}

func provideRedisConn(cfg *config.Config) redis.Provider {
	return redis.NewRedisConnection(cfg.Redis)
}

func provideObjectCacher[T any](p redis.Provider) *redis.ObjectCacher[T] {
	return redis.NewObjectCacher[T](p)
}

func provideOTPRepo(oc *redis.ObjectCacher[types.OTP]) repository.OtpRepo {
	return repository.NewOtpRepo(oc)
}

func provideOTPService(repo repository.OtpRepo) *service.OTPService {
	return service.NewOTPService(repo)
}
