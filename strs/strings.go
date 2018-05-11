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
}`
	UtilsFile string = `package utils

import "log"

// CheckErr verify if exists an error in the var err
func CheckErr(err error, msg string) {
	if err != nil {
		log.Println(err, msg)
	}
}`

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
	utils.CheckErr(err, "No se puede abrir una conexión")

	return db
}`
)

func GetVar(name string) string {
	switch name {
		case "MainFile":
			return MainFile
		case "HomeActionsFile":
			return HomeActions
		case "RoutesFile":
			return RoutesFile
		case "UtilsFile":
			return UtilsFile
		case "DBFile":
			return DBFile
	}
	return ""
}
