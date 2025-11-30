package main

import (
	"ibnufth/backend-api/config"
	"ibnufth/backend-api/database"
	"ibnufth/backend-api/routes" 
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize database
	database.InitDB()

	// setup router
	router := routes.SetupRouter()

	// menjalankan server pada port yang ditentukan
	port := config.GetEnv("APP_PORT", "3000")
	router.Run(":" + port)
}
