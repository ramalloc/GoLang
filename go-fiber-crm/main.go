package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/ramalloc/go-fiber-crm/database"
	"github.com/ramalloc/go-fiber-crm/lead"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/v1/leads", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
	app.Delete("/api/v1/leads", lead.DeleteLeads)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("Database Connection Failed...")
	}
	fmt.Println("Connection Opened With Databse...")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Databse Migrated...")
}

func main() {
	app := fiber.New()
	initDatabase()
	SetupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()
}
