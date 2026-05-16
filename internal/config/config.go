package config

type Config struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string

	CORSAllowedOrigins string
	CORSAllowedHeaders string
	CORSAllowedMethods string
	CORSMaxAgeSeconds  int

	RedisHost string
	RedisPort string
}

func LoadConfig() *Config {
	cfg := &Config{
		DBHost:     getEnv("DB_HOST", "db"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBName:     getEnv("DB_NAME", "civil"),
		DBUser:     getEnv("DB_USER", "civil_user"),
		DBPassword: getEnv("DB_PASSWORD", ""),

		CORSAllowedOrigins: getEnv("CORS_ALLOWED_ORIGINS", "*"),
		CORSAllowedHeaders: getEnv("CORS_ALLOWED_HEADERS", "*"),
		CORSAllowedMethods: getEnv("CORS_ALLOWED_METHODS", "*"),
		CORSMaxAgeSeconds:  getEnvAsInt("CORS_MAX_AGE_SECONDS", 300),

		RedisHost: getEnv("REDIS_HOST", "redis"),
		RedisPort: getEnv("REDIS_PORT", "6379"),
	}

	return cfg
}
