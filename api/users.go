package api

import (
	"context"
	"fmt"

	_ "database/sql"

	db "github.com/Richd0tcom/BE-Stage-1/db/sqlc"
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/google/uuid"
)

type CreateUseReq struct{
	Name string `json:"name"`
}

type UpdateUseReq struct{
	NewName string `json:"new_name"`
}


func (s *Server) CreateUser(ctx *fiber.Ctx) error {
	fmt.Println("I ran here")
	bdy := CreateUseReq{}
	if err:= ctx.BodyParser(&bdy); err != nil {
		fmt.Println("The error was here")
		return err
	}

	user, err := s.store.CreateUser(context.Background(), bdy.Name)

	if err != nil {
		return err
	}

	return ctx.JSON(user)

}

func (s *Server) GetUser(ctx *fiber.Ctx) error {
	someNam:= ctx.Query("name")
	if someNam == "" {
		user_id:= ctx.Params("user_id")
		parsedId, _:= uuid.Parse(user_id)
		user, err := s.store.GetUserById(context.Background(), parsedId)
		if err != nil {
			return err
		}
		return ctx.JSON(user)
	}
	user, err := s.store.GetUserByName(context.Background(), someNam)
	if err != nil {
		return err
	}

	return ctx.JSON(user)

}


func (s *Server) UpdateUser(ctx *fiber.Ctx) error {
	bdy := UpdateUseReq{}
	if err:= ctx.BodyParser(&bdy); err != nil {
		return err
	}

		someNam:= ctx.Query("name")
	if someNam == "" {
		user_id:= ctx.Params("user_id")
		parsedId, _:= uuid.Parse(user_id)
		user, err := s.store.UpdateUserById(context.Background(), db.UpdateUserByIdParams{
			ID: parsedId,
			Name: bdy.NewName,
		})
		if err != nil {
			return err
		}
		return ctx.JSON(user)
	}
	user, err := s.store.UpdateUserByName(context.Background(), db.UpdateUserByNameParams{
		Name: someNam,
		Name_2: bdy.NewName,
	})
	if err != nil {
		return err
	}

	return ctx.JSON(user)

}

func (s *Server) DeleteUser(ctx *fiber.Ctx) error {

	type res struct {
		message string
	}

		someNam:= ctx.Query("name")
	if someNam == "" {
		user_id:= ctx.Params("user_id")
		parsedId, _:= uuid.Parse(user_id)
		err := s.store.DeleteUserById(context.Background(), parsedId)
		if err != nil {
			return err
		}
		return ctx.JSON(res{message: "user deleted",})
	}
	err := s.store.DeleteUserByName(context.Background(), someNam)
	if err != nil {
		return err
	}

	return ctx.JSON(res{message: "user deleted",})

}