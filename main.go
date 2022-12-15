package main

import (
	"log"
	"os"
	"shopular/ent"
	"shopular/handlers"
	"shopular/services"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	app := fiber.New()
	client, err := ent.Open("mysql", os.Getenv("MYSQL_PATH"))
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// if err := client.Schema.Create(context.Background()); err != nil {
	// 	log.Fatalf("failed creating schema resources: %v", err)
	// }
	sp := services.NewServiceProvider(client)
	handlers.SetUpHandlers(app, &sp)
	app.Listen(":3000")
}
