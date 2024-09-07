package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/gofiber/fiber"
	"github.com/raghulchandramouli/go-fiber-crm-basic/lead"
	"github.com/raghulchandramouli/go-fiber-crm-basic/database"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead",lead.GetLeads)
	app.Get("/api/v1/lead:id",lead.GetLead)
	app.Post("/api/v1/lead",lead.NewLead)
	app.Delete("/api/v1/lead:id",lead.DeleteLead)

}



func initDatabase(){
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")

	// If database fails to open, or connection is closed
	if err != nil {
		panic("failed to connect to database")
	}

	// AutoMigrate the schema of the model to the database
	fmt.Println("Database connected successfully")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database successfully")

}
func main() {
	app := fiber.New()
	initDatabase()

	setupRoutes(app)

	// Port Starter
	app.Listen(3000)

	// slow shutdown of the database
	defer database.DBConn.Close()
}