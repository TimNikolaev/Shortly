package configs

type Config struct {
	DBConfig
	AuthConfig
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
}

type AuthConfig struct {
	Secret string
}
