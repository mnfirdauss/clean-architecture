package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	userEntity "github.com/mnfirdauss/clean-architecture/entity/user"
)

func (h *UserHandler) CreateUser(c echo.Context) error {
	data := map[string]interface{}{
		"message": "fail",
	}
	var user userEntity.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, data)

	}

	err = h.userUC.CreateUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, data)
	}

	return c.JSON(http.StatusOK, data)
}
