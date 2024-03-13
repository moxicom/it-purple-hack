package config

import "os"

type DBInfo struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresName     string
	PostgresSSLMode  string
}

type RedisInfo struct {
	RedisAddr     string
	RedisPassword string
	RedisDB       string
}

func ReadDbInfo() DBInfo {
	return DBInfo{
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresName:     os.Getenv("POSTGRES_NAME"),
		PostgresSSLMode:  os.Getenv("POSTGRES_SSLMODE"),
	}
}

func ReadRedisInfo() RedisInfo {
	return RedisInfo{
		RedisAddr:     os.Getenv("REDIS_ADDR"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisDB:       os.Getenv("REDIS_DB"),
	}
}
