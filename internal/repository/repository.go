package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/moxicom/it-purple-hack/config"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	db      *sql.DB
	cacheDb *redis.Client
}

func NewRepository(db *sql.DB, cacheDb *redis.Client) *Repository {
	return &Repository{db, cacheDb}
}

func NewPostgres(config config.DBInfo) *sql.DB {
	fmt.Println("Connecting to database...")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.PostgresHost,
		config.PostgresPort,
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresName,
		config.PostgresSSLMode,
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Database connected")
	return db
}

func NewRedis(config config.RedisInfo) *redis.Client {
	fmt.Println("Connecting to redis...")
	redisDBint, err := strconv.Atoi(config.RedisDB)
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       redisDBint,
	})
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Connected to redis")
	return client
}
