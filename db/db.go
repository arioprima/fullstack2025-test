package db

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var DB *sql.DB
var RDB *redis.Client

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var err error
	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s",
		viper.GetString("postgres.user"),
		viper.GetString("postgres.dbname"),
		viper.GetString("postgres.password"),
	)
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	RDB = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})

	ctx := context.Background()
	_, err = RDB.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
}
