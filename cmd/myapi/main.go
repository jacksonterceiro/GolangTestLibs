package main

// import (
// 	"bytes"
// 	"fmt"
// 	"log"
// 	"myapi/internal/pkg/health-check"
// 	"myapi/internal/pkg/user"

// 	"github.com/labstack/echo"
// )

// func main() {
// 	//redis.ExampleClient()

// 	e := echo.New()

// 	e.GET("/health", health.Health)
// 	e.GET("/users/:id", user.GetUser)

// 	// e.POST("/users", saveUser)
// 	// e.PUT("/users/:id", updateUser)
// 	// e.DELETE("/users/:id", deleteUser)

// 	e.Logger.Fatal(e.Start(":1323"))

// 	var (
// 		buf    bytes.Buffer
// 		logger = log.New(&buf, "logger: ", log.Lshortfile)
// 	)

// 	logger.Print("Hello, log file!")

// 	fmt.Print(&buf)
// }

import (
	"myapi/internal/pkg/health-check"
	levelog "myapi/internal/pkg/level-log"
	"myapi/internal/pkg/user"
	"os"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func main() {

	_ = os.MkdirAll("/log", 0777)
	file, err := os.OpenFile("/log/logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.SetFormatter(&log.JSONFormatter{})

	log.Info("Starting api...")

	e := echo.New()

	e.GET("/health", health.Health)
	e.GET("/users/:id", user.GetUser)

	e.GET("/levelog/info", levelog.LogInfo)
	e.GET("/levelog/debug", levelog.LogDebug)
	e.GET("/levelog/warn", levelog.LogWarn)
	e.GET("/levelog/err", levelog.LogError)

	// e.POST("/users", saveUser)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}
