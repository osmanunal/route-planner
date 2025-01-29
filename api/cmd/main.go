package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"route-planner/api/internal/middleware"
	"route-planner/api/internal/router"
	"route-planner/pkg/config"
	"route-planner/pkg/database"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	fiberApp := fiber.New(fiber.Config{
		ReadTimeout:  time.Duration(30) * time.Second,
		WriteTimeout: time.Duration(30) * time.Second,
		IdleTimeout:  time.Duration(120) * time.Second,
	})
	fiberApp.Use(middleware.RateLimiter())
	fiberApp.Use(cors.New())
	fiberApp.Use(requestid.New(requestid.Config{
		Header:     fiber.HeaderXRequestID,
		ContextKey: "requestid",
	}))

	DB := database.ConnectDB(cfg.DBConfig)
	router.Setup(fiberApp, DB)
	log.Println("http server başlatılıyor...")
	go func() {
		err = fiberApp.Listen(fmt.Sprintf(":%v", 3000))
		if err != nil {
			panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	log.Println("Gracefully shutting down...")
	err = fiberApp.Shutdown()
	if err != nil {
		log.Println("FiberApp shutdown", err)
	}
	err = DB.Close()
	if err != nil {
		log.Println("DB shutdown", err)
	}
	log.Println("Server stopped")
}
