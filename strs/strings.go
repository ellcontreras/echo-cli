package strs

var (
	MainFile string = `package main

import (
	"appName/routes"

	"github.com/labstack/echo"
)

var (
	//Server is the server of the app
	Server *echo.Echo
)

func main() {
	Server = echo.New()

	//Init the routes
	routes.InitRoutes(Server)

	Server.Logger.Fatal(Server.Start(":8080"))
}`

	RoutesFile string = `package routes

import (
	"appName/actions"

	"github.com/labstack/echo"
)

//InitRoutes the routes for the server
func InitRoutes(server *echo.Echo) {
	//Routes...
}`
)
