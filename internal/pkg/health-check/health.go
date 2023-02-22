package health

import (
	"net/http"

	"github.com/labstack/echo"
)

type Response struct {
	Message string `json:"message"`
}

func Health(c echo.Context) error {
	return c.JSON(http.StatusOK, Response{Message: "Serviço UP!"})
}
