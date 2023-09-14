package api

import (
	"database/sql"

	dst "github.com/Richd0tcom/BE-Stage-1/db/sqlc"
	"github.com/gofiber/fiber/v2"
)

type Store struct {
	*dst.Queries
	Gb *sql.DB
}

type Server struct {
	store Store
	Srv *fiber.App
}
func NewStore(db *sql.DB) Store {
	return Store{
		Gb:      db,
		Queries: dst.New(db),
	}
}

func NewServer(store Store) *Server {
	sv:= &Server{store: store}

	sv.Srv = fiber.New()
	return sv
}