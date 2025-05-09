package app


import (
	"github.com/mojtabamovahedi/otp/config"
	"github.com/mojtabamovahedi/otp/pkg/logger"
	redisConn "github.com/mojtabamovahedi/otp/pkg/redis"
)

type App struct {
	cfg config.Config
	logger logger.Logger
	redisConn redisConn.Provider
}


func NewApp(cfg config.Config, logger logger.Logger, redisConnection redisConn.Provider) *App {
	return &App{
		cfg: cfg,
		logger: logger,
		redisConn: redisConnection,
	}
}

func (a *App) Config() config.Config {
	return a.cfg
}

func (a *App) Logger() logger.Logger {
	return a.logger
}

func (a *App) RedisConnection() redisConn.Provider {
	return a.redisConn
}