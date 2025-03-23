package main

import (
	"fmt"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/config"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/database"
	"github.com/spitfireooo/form-constructor-server-v2/pkg/router"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatalln("Error in loading env file", err)
	}

	if err := config.ConfigInit(); err != nil {
		PORT = ":8060"
		log.Fatalln("Error in configuration init", err)
	} else {
		PORT = fmt.Sprintf(":%v", viper.GetString("http.port"))
	}

	if err := database.DatabaseInit(database.ConnectConfig{
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: viper.GetString("db.database"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		SSLMode:  viper.GetString("db.ssl_mode"),
	}); err != nil {
		log.Fatalln("Error in database connection", err)
	} else {
		log.Println("Database connected successfully")
	}
}

var PORT string

func main() {
	app := fiber.New(fiber.Config{
		// Prefork: true,
		AppName: "Form Constructor",
	})

	app.Static("/static", "./static")

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(swagger.New(swagger.Config{
		BasePath: "/docs",
		FilePath: "./docs/swagger.json",
		Path:     "docs",
	}))

	router.Router(app)

	if err := app.Listen(PORT); err != nil {
		log.Fatalln("Error in server started", err)
		return
	}
}
