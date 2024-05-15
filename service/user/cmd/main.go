package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"os"
	"time"
	"user/handler"
)

func main() {
	e := echo.New()

	var (
		redisHost = os.Getenv("REDIS_HOST")
		redisPort = os.Getenv("REDIS_PORT")
	)

	redisClient := redis.NewClient(
		&redis.Options{
			Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
		},
	)

	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()
	if err := redisClient.Ping(ctx).Err(); err != nil {
		e.Logger.Fatalf("failed to connect to redis: %v", err)
	}

	e.GET("/users/:id", handler.GetUserHandler)
	e.GET("/count", func(c echo.Context) error {
		return handler.GetCountHandler(c, redisClient)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
