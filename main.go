package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"log"
	"strings"
	"echo-cli/strs"
	"os"
)

//checkErr check if exists an error in the param err
func checkErr(err error) {
	if err != nil {
		log.Println("Error:", err)
	}
}

//checkNull check if a string is null or not
func checkNull(newApp *string) bool {
	log.Print("Value:", *newApp)
	return (*newApp != "" || *newApp != " ")
}

func main() {
	fmt.Println("Echo Framework CLI\n")

	newApp := flag.String("new", "nameApp", "Generate de folder structure of an app")
	//generate := flag.String("generate", "", "Generate something such as controller, model, route")

	flag.Parse()

	if checkNull(newApp) {
		strParse := strings.Replace(strs.MainFile, "appName", *newApp,-1)
		log.Print(strParse)
		d1 := []byte(strParse)
		os.MkdirAll(*newApp, os.ModePerm)
		err := ioutil.WriteFile(*newApp+"/main.go", d1, 0644)
		checkErr(err)
	}
}
