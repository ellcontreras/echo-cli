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

	//Make migrations
	//db := DB.Open()
	//db.AutoMigrate(&models.Product{})

	Server.Logger.Fatal(Server.Start(":8080"))
}

`

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
}

`
	RoutesFile string = `package routes

import (
	"appName/actions"

	"github.com/labstack/echo"
)

//InitRoutes the routes for the server
func InitRoutes(server *echo.Echo) {
	//User routes
	server.GET("/api/home/get", actions.HomeGet)
	server.POST("/api/home/post", actions.HomePost)
	server.PUT("/api/home/put", actions.HomePut)
	server.DELETE("/api/home/delete", actions.HomeDelete)
}

`
	UtilsFile string = `package utils

import "log"

// CheckErr verify if exists an error in the var err
func CheckErr(err error, msg string) {
	if err != nil {
		log.Println(err, msg)
	}
}

`

	DBFile string = `package DB

import (
	"appName/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db  *gorm.DB
	err error
)

//Open return a dastabase instance
func Open() *gorm.DB {
	db, err = gorm.Open("mysql", "root:root@/appName")
	utils.CheckErr(err, "Connection can not be opened")

	return db
}

`
	Action = `package actions

import (
	"net/http"
	"github.com/labstack/echo"
)

//appNameGet ...
func appNameGet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Get Method")
}

//appNamePost ...
func appNamePost(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Post Method")
}

//appNamePut ...
func appNamePut(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Put Method")
}

//Delete ...
func appNameDelete(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Delete Method")
}

	`
)

var Paths map[string]string = make(map[string]string)

func Init() {
	//Main File
	Paths["main.go"] = MainFile
	//Routes
	Paths["routes/Routes.go"] = RoutesFile
	//Controller
	Paths["actions/HomeActions.go"] = HomeActions
	//Utils
	Paths["utils/utils.go"] = UtilsFile
	//DB
	Paths["db/db.go"] = DBFile
	//Actions
	Paths["actions/action.go"] = Action
}
