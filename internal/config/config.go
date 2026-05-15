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
}

func Load() *Config {
	cfg := &Config{
		DBHost:     getEnv("DB_HOST", ""),
		DBPort:     getEnv("DB_PORT", ""),
		DBName:     getEnv("DB_NAME", ""),
		DBUser:     getEnv("DB_USER", ""),
		DBPassword: getEnv("DB_PASSWORD", ""),

		CORSAllowedOrigins: getEnv("CORS_ALLOWED_ORIGINS", ""),
		CORSAllowedHeaders: getEnv("CORS_ALLOWED_HEADERS", ""),
		CORSAllowedMethods: getEnv("CORS_ALLOWED_METHODS", ""),
		CORSMaxAgeSeconds:  getEnvAsInt("CORS_MAX_AGE_SECONDS", 300),
	}

	return cfg
}
