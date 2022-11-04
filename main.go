package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)
type User struct {
	SlackUsername string `json:"slackUsername"`
	Backend bool `json:"backend"`
	Age int32 `json:"age"`
	Bio string `json:"bio"`
}

type OperationRequest struct {
	OperationType string `json:"operation_type"`
	X int `json:"x"`
	Y int `json:"y"`
}

type OperationResponse struct {
	SlackUsername string `json:"slackUsername"`
	Result int `json:"result"`
	OperationType string `json:"operation_type"`
	
}

func main(){
	fmt.Print("Hello, world!")

	//create a new app.
	app := fiber.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app.Use(cors.New(cors.Config{AllowOrigins: "*"}))

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

	app.Post("/operation", func(fiberContext *fiber.Ctx) error {
		sum := OperationRequest{}

		if err := fiberContext.BodyParser(&sum); err != nil {
			return err
		}
		res, opType := checkOperation(strings.ToLower(sum.OperationType), sum.X, sum.Y)
		result := OperationResponse{
			SlackUsername: "Richdotcom",
			OperationType: opType,
			Result: res,
		}
		return fiberContext.JSON(result)
	})
	log.Fatal(app.Listen(`:`+port))
}

func checkOperation(opType string, x int, y int) (int, string){
	var add = regexp.MustCompile(`add|sum|summation|plus|addition`)
	var subtract = regexp.MustCompile(`subtract|minus|subtraction`)
	var multiply  = regexp.MustCompile(`multiply|multiplication|product`)
	var re int
	var operationType string
	switch {
		case opType == "addition":
			re= x + y
			operationType = "addition"
			break
		case opType == "subtraction":
			re= x - y
			operationType = "subtraction"
			break
		case opType == "multiplication":
			re= x * y
			operationType = "multiplication"
			break
		case add.MatchString(opType):
			re= x + y
			operationType = "addition"
			break
		case subtract.MatchString(opType):
			re= x - y
			operationType = "subtraction"
			break
		case multiply.MatchString(opType):
			re= x * y
			operationType = "multiplication"
			break

	}

	return re, operationType
}