package config

import (
    "os"
)

type Config struct {
    ServerPort  string
    DBHost      string
    DBUser      string
    DBPassword  string
    DBName      string
    RedisAddr   string
    RedisPassword string
    JWTSecret   string
}

func InitConfig() *Config {
    return &Config{
        ServerPort:  getEnv("SERVER_PORT", "5001"),
        DBHost:      getEnv("DB_HOST", "localhost"),
        DBUser:      getEnv("DB_USER", "postgres"),
        DBPassword:  getEnv("DB_PASSWORD", ""),
        DBName:      getEnv("DB_NAME", "testdb"),
        RedisAddr:   getEnv("REDIS_ADDR", "localhost:6379"),
        RedisPassword: getEnv("REDIS_PASSWORD", ""),
        JWTSecret:   getEnv("JWT_SECRET", "your_jwt_secret"),
    }
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}
