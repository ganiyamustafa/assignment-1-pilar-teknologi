package main

import (
	"log"
	"os"

	"github.com/ganiyamustafa/assignment-1-pilar-teknologi/db/connections"
	"github.com/ganiyamustafa/assignment-1-pilar-teknologi/db/migrations"
	"github.com/ganiyamustafa/assignment-1-pilar-teknologi/db/seeders"
	"github.com/ganiyamustafa/assignment-1-pilar-teknologi/internal/routes"
	"github.com/ganiyamustafa/assignment-1-pilar-teknologi/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	// Run program if no flag is given
	if len(os.Args) < 2 {
		app := gin.Default()
		app.Use(cors.New(cors.Config{
			AllowAllOrigins: true,
			AllowMethods:    []string{"*"},
			AllowHeaders:    []string{"*"},
		}))

		// init the main handler that will be used throughout the application
		handler := utils.Handler{}
		handler.Validator = validator.New()

		if err := connections.ConnectSqlite(); err != nil {
			log.Fatal("Failed to connect to sqlite: " + err.Error())
		}
		handler.SQLite = connections.SQLite

		// Register router to app with handler
		routes.ChatRoutes(app, &handler)

		if err := app.Run("127.0.0.1:8000"); err != nil {
			return
		}
	}

	command := os.Args[1]
	switch command {
	case "migrate":
		if err := connections.ConnectSqlite(); err != nil {
			log.Fatal("Failed to connect to postgres")
		}

		migrations.Migrate(connections.SQLite)
	case "rollback":
		if err := connections.ConnectSqlite(); err != nil {
			log.Fatal("Failed to connect to postgres")
		}

		migrations.Rollback(connections.SQLite)
	case "seed":
		if err := connections.ConnectSqlite(); err != nil {
			log.Fatal("Failed to connect to postgres")
		}

		seeders.Seed(connections.SQLite)
	case "wipe":
		// Delete data from all table that have a seeder
		if err := connections.ConnectSqlite(); err != nil {
			log.Fatal("Failed to connect to postgres")
		}

		seeders.Wipe(connections.SQLite)
	default:
		log.Fatal("Unknown command flag")
	}
}
