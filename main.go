package main

import (
	"fmt"
	"log"
	"os"
	"time"

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
	SlackUsername string `json:"slack_name"`
	Result int `json:"result"`
	OperationType string `json:"operation_type"`
	
}

type HngXStageOneResponse struct {
	SlackUsername string `json:"slack_name"`
	CurrentDay string `json:"current_day"`
	UtcTime time.Time `json:"utc_time"`
	Track string `json:"track"`
	GithubFileUrl string `json:"github_file_url"`
	GithubRepoUrl string `json:"github_repo_url"`
	StatusCode int `json:"status_code"`
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

	

	app.Get("/api", func(fiberContext *fiber.Ctx) error {
	
		track := fiberContext.Query("track")
		slack_user := fiberContext.Query("slack_name")
		result := HngXStageOneResponse{

			SlackUsername: slack_user,
			CurrentDay: time.Weekday.String(time.Now().Weekday()),
			UtcTime: time.Now(),
			Track: track,
			GithubFileUrl: "https://github.com/Richd0tcom/BE-Stage-1/blob/main/main.go",
			GithubRepoUrl: "https://github.com/Richd0tcom/BE-Stage-1",
			StatusCode: 200,
		}
		return fiberContext.JSON(result)
	})
	log.Fatal(app.Listen(`:`+port))
}

