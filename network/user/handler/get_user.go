package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"user/internal/response"
	"user/model"
)

type GetUserResponse struct {
	User *model.User `json:"user,omitempty"`
}

func GetUserHandler(c echo.Context) error {
	// User ID from path `users/:id`
	userIDStr := c.Param("id")

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &response.ErrorResponse{
			Message: "invalid user id",
		})

	}

	user := getUserByID(userID)
	if user == nil {
		return c.JSON(http.StatusNotFound, &response.ErrorResponse{
			Message: "user not found",
		})

	}

	return c.JSON(http.StatusOK, &GetUserResponse{
		User: user,
	})
}

func getUserByID(userID int64) *model.User {
	users := []*model.User{
		{1, "Alice"},
		{2, "Bob"},
		{3, "Charlie"},
		{4, "David"},
		{5, "Eve"},
	}

	userById := make(map[int64]*model.User, len(users))
	for _, user := range users {
		userById[user.ID] = user
	}

	return userById[userID]
}
