package config

type PSQL struct {
	User string
	DBname string
	Password string
	SSlmode string
}

type Server struct {
	Port string
}

type Config struct {
	PSQL PSQL
	Server Server
}

func NewConfig() *Config {
	var cfg Config

	cfg.PSQL.DBname = "instagram"
	cfg.PSQL.User = "postgres"
	cfg.PSQL.Password = "salom"
	cfg.PSQL.SSlmode = "disable"

	cfg.Server.Port = ":8080"

	return &cfg
}