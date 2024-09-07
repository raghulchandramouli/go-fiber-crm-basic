package main

import (
	"github.com/gofiber/fiber"
)

func setupRoutes(spp *fiber.App) {
	app.Get(GetLeads)
	app.Get(GetLead)
	app.Post(NewLead)
	app.Delete(DeleteLead)

}

func initDatabase(){

	
}
func main() {
	app := fiber.New()
	setupRoutes(app)

	// Port Starter
	app.Listen(3000)
}