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

	HomeActions string = `package actions

import (
	"net/http"
	"github.com/labstack/echo"
)

//HomeGet ...
func HomeGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Get Method")
}

//HomePost ...
func HomePost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Post Method")
}

//HomePut ...
func HomePut(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Put Method")
}

//HomeDelete ...
func HomeDelete(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Delete Method")
}`
)

func GetVar(name string) string {
	//switch name {
	//	case "MainFile":
	//		return name
	//		break
	//	case "HomeActions":
	//		return HomeActions
	//		break
	//}
	return name
}
