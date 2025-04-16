package main

import (
	"fmt"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
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
		log.Fatalln("Error in configuration init", err)
	} else {
		PORT = viper.GetInt("http.port")
	}

	cmd := config.CobraInit(config.CobraConfig{
		Use:   "form constructor",
		Short: "Server for form constructor",
		Long:  "Server for form constructor",
	})
	cmd.PersistentFlags().IntVarP(&PORT, "port", "p", 8060, "Port for starting")
	cmd.PersistentFlags().StringVarP(&DB_CON, "db-con", "d", "uri", "Database mode connection")
	if err := cmd.Execute(); err != nil {
		log.Fatalln("Error in cobra init", err)
	}

	var db_err error
	if DB_CON == "conf" {
		db_err = database.DatabaseInit(database.ConnectConfig{
			Username: viper.GetString("db.username"),
			Password: os.Getenv("DB_PASSWORD_LOCAL"),
			Database: viper.GetString("db.database"),
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			SSLMode:  viper.GetString("db.ssl_mode"),
		})
	} else if DB_CON == "url" || DB_CON == "uri" {
		db_err = database.DatabaseInitByURL(os.Getenv("DATABASE_URL"))
	}
	if db_err != nil {
		log.Fatal("Error in database connection", db_err)
	}
}

var (
	PORT   int
	DB_CON string
)

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
	app.Use(recover.New())

	router.Router(app)

	if err := app.Listen(fmt.Sprintf(":%v", PORT)); err != nil {
		log.Fatalln("Error in server started", err)
		return
	}
}
