package user

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

type User struct {
	IdUser int    `json:"idUser,omitempty"`
	Name   string `json:"name,omitempty"`
}

// e.GET("/users/:id", getUser)
func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	log.Info("Call GetUser Method 2...")

	return c.JSON(http.StatusOK, User{IdUser: id})
}

// e.POST("/users", saveUser)
func saveUser(c echo.Context) error {
	return nil
}

// e.PUT("/users/:id", updateUser)
func updateUser(c echo.Context) error {
	return nil
}

// e.DELETE("/users/:id", deleteUser)
func deleteUser(c echo.Context) error {
	return nil
}
