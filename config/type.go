package config

type Config struct {
	Server ServerConfig `yaml:"server"`
	Redis  RedisConfig  `yaml:"redis"`
}

type ServerConfig struct {
	HttpPort uint `yaml:"http_port"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     uint   `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}
