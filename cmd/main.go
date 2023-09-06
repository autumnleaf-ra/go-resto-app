package main

import (
	"github.com/autumnleaf-ra/go-resto-app/internal/database"
	"github.com/autumnleaf-ra/go-resto-app/internal/delivery/rest"
	mRepo "github.com/autumnleaf-ra/go-resto-app/internal/repository/menu"
	rUsecase "github.com/autumnleaf-ra/go-resto-app/internal/usecase/resto"
	"github.com/labstack/echo/v4"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password=postgres dbname=go_resto_app sslmode=disable"
)

func main() {
	e := echo.New()

	db := database.GetDB(dbAddress)
	menuRepo := mRepo.GetRepository(db)
	restoUsecase := rUsecase.GetUsecase(menuRepo)

	h := rest.NewHandler(restoUsecase)

	rest.LoadRoutes(e, h)

	// localhost:14045/menu/food
	e.Logger.Fatal(e.Start(":14045"))
}
