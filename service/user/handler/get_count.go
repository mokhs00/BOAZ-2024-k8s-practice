package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"user/internal/response"
)

type GetCountResponse struct {
	Count int64 `json:"count"`
}

func GetCountHandler(c echo.Context, redisClient *redis.Client) error {
	ctx := c.Request().Context()

	count := redisClient.Incr(ctx, "count")
	if count.Err() != nil {
		return c.JSON(500, &response.ErrorResponse{
			Message: fmt.Sprintf("internal server error: '%s'", count.Err()),
		})
	}

	return c.JSON(200, &GetCountResponse{
		Count: count.Val(),
	})
}
