package main

import (
	"fmt"
	"log"
	"os"
	// "time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"database/sql"

	"github.com/Richd0tcom/BE-Stage-1/api"
	// db "github.com/Richd0tcom/BE-Stage-1/db/sqlc"
	"github.com/Richd0tcom/BE-Stage-1/utils"

	_ "github.com/lib/pq"
)
type User struct {
	SlackUsername string `json:"slackUsername"`
	Backend bool `json:"backend"`
	Age int32 `json:"age"`
	Bio string `json:"bio"`
}


type HngXStageOneResponse struct {
	SlackUsername string `json:"slack_name"`
	CurrentDay string `json:"current_day"`
	UtcTime string `json:"utc_time"`
	Track string `json:"track"`
	GithubFileUrl string `json:"github_file_url"`
	GithubRepoUrl string `json:"github_repo_url"`
	StatusCode int `json:"status_code"`
}
func main(){
	fmt.Print("Hello, world!")


	config, err:= utils.LoadConfig(".")
	if err != nil {
		log.Fatal("could not read configs ", err)
	}

	conn, err := sql.Open(config.DbDriver, config.DbUri)
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	store:= api.NewStore(conn)

	srv:=api.NewServer(store)

	
	app:= srv.Srv
	//create a new app.
	// app := fiber.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app.Use(cors.New(cors.Config{AllowOrigins: "*"}))

	

	app.Get("/healthcheck", func(fiberContext *fiber.Ctx) error {
		return fiberContext.SendString("OK")
	})

	app.Post("/api/users", srv.CreateUser)
	app.Get("/api/users/:user_id", srv.GetUser)
	app.Put("/api/users/:user_id", srv.UpdateUser)
	app.Delete("/api/users/:user_id", srv.DeleteUser)
	

	// app.Get("/api", func(fiberContext *fiber.Ctx) error {
	
	// 	track := fiberContext.Query("track")
	// 	slack_user := fiberContext.Query("slack_name")
	// 	result := HngXStageOneResponse{

	// 		SlackUsername: slack_user,
	// 		CurrentDay: time.Weekday.String(time.Now().Weekday()),
	// 		UtcTime: time.Now().UTC().Format(time.RFC3339),
	// 		Track: track,
	// 		GithubFileUrl: "https://github.com/Richd0tcom/BE-Stage-1/blob/main/main.go",
	// 		GithubRepoUrl: "https://github.com/Richd0tcom/BE-Stage-1",
	// 		StatusCode: 200,
	// 	}
	// 	return fiberContext.JSON(result)
	// })

	log.Fatal(app.Listen(`:`+port))
}

