package main

import (
	"aquaculture/database"
	"aquaculture/routes"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDB()

	e := echo.New()

	routes.SetupRoutes(e)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	e.Logger.Fatal(e.Start(port))
}
