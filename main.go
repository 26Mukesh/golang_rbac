package main

// load required packages
import (
	"fmt"
	"jwt-go-rbac/config"
	"jwt-go-rbac/database"
	"jwt-go-rbac/model"
	"jwt-go-rbac/routes"
	"os"

	"github.com/gin-gonic/gin"
)

// main function
func main() {
	// load environment file
	config.LoadEnv()
	// load database configuration and connection
	loadDatabase()
	// start the server
	serveApplication()
}

// run database migrations and add seed data
func loadDatabase() {
	database.InitDb()
	database.Db.AutoMigrate(&model.Role{})
	database.Db.AutoMigrate(&model.User{})
	database.Db.AutoMigrate(&model.Room{})
	database.Db.AutoMigrate(&model.Booking{})
	seedData()
}

// load seed data into the database
func seedData() {
	var roles = []model.Role{{Name: "admin", Description: "Administrator role"}, {Name: "customer", Description: "Authenticated customer role"}, {Name: "visitor", Description: "Unauthenticated customer role"}}
	var user = []model.User{{Username: os.Getenv("ADMIN_USERNAME"), Email: os.Getenv("ADMIN_EMAIL"), Password: os.Getenv("ADMIN_PASSWORD"), RoleID: 1, EmailVerified: true}}
	database.Db.Save(&roles)
	database.Db.Save(&user)
}

// start the server on port 8000
func serveApplication() {
	router := gin.Default()
	// Initialize routes
	routes.Routes(router)

	router.Run(":" + os.Getenv("PORT"))
	fmt.Println("Server running on port 8000")
}
