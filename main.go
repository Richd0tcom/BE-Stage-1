package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)
type User struct {
	SlackUsername string `json:"slackUsername"`
	Backend bool `json:"backend"`
	Age int32 `json:"age"`
	Bio string `json:"bio"`
}
func main(){
	fmt.Print("Hello, world!")

	//create a new app.
	app := fiber.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app.Get("/healthcheck", func(fiberContext *fiber.Ctx) error {
		return fiberContext.SendString("OK")
	})
	var user User

	app.Get("/user", func(fiberContext *fiber.Ctx) error {
		user.SlackUsername = "Richdotcom"
		user.Backend = true
		user.Age = 23
		user.Bio = "Creative, detail-oriented software developer"

		return fiberContext.JSON(user)
	})
	log.Fatal(app.Listen(`:`+port))
}