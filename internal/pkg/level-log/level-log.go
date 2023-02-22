package levelog

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func LogInfo(c echo.Context) error {
	log.Info("logging in level info...")

	return c.JSON(http.StatusOK, "Log info OK")
}

func LogDebug(c echo.Context) error {
	log.Debug("logging in level debug...")

	return c.JSON(http.StatusOK, "Log debug OK")
}

func LogWarn(c echo.Context) error {
	log.Warn("logging in level warn...")

	return c.JSON(http.StatusOK, "Log warn OK")
}

func LogError(c echo.Context) error {
	log.Error("logging in level error...")

	return c.JSON(http.StatusOK, "Log error OK")
}
