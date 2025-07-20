package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tiedsandi/fleetify-backend-fachransandi/config"
	"github.com/tiedsandi/fleetify-backend-fachransandi/routes"
	"github.com/tiedsandi/fleetify-backend-fachransandi/seeds"
)

func main() {
	config.InitDB()
	// config.ResetDB()
	config.Migration()
	seeds.AddSeeder()

	server := gin.Default()

	server.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"http://localhost:5173"},
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.RegisterRoutes(server)
	server.Run()
}
