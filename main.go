package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"log"
	"strings"
	"os"
	"echo-cli/strs"
)

type FileParse struct {
	name		string
	appName		string
	subFolder	string
	fileName	string
	strName		string
}

func (f *FileParse) Parse(){
	strParse := strings.Replace(strs.GetVar(f.strName), "appName", f.appName,-1)
	d1 := []byte(strParse)
	os.MkdirAll(f.appName, os.ModePerm)

	if checkNil(f.subFolder) {
		os.MkdirAll(f.appName + "/" + f.subFolder, os.ModePerm)
		err := ioutil.WriteFile(f.appName+"/" + f.subFolder + f.fileName + ".go", d1, 0644)
		checkErr(err)
	} else {
		err := ioutil.WriteFile(f.appName+"/" + f.fileName + ".go", d1, 0644)
		checkErr(err)
	}
}

//checkErr check if exists an error in the param err
func checkErr(err error) {
	if err != nil {
		log.Println("Error:", err)
	}
}

//checkNull check if a string is null or not
func checkNull(newApp *string) bool {
	log.Print(*newApp != "")
	return *newApp == ""
}

//checkNil check if non pointer of a string is null
func checkNil(a string) bool {
	return a == ""
}

func main() {
	fmt.Println("Echo Framework CLI")

	newApp := flag.String("new", "nameApp", "Generate de folder structure of an app")
	//generate := flag.String("generate", "", "Generate something such as controller, model, route")
	//flag.PrintDefaults()
	flag.Parse()

	if checkNull(newApp) {
		fMain := FileParse{"Main File", *newApp, "","main", "MainFile"}
		fMain.Parse()

		fActions := FileParse{"Actions folder and file", *newApp, "actions/", "HomeActions", "HomeActionsFile"}
		fActions.Parse()

		fRoutes := FileParse{"Routes folder and file", *newApp, "routes/", "Routes", "RoutesFile"}
		fRoutes.Parse()

		fUtils := FileParse{"Util function such as checkErr", *newApp, "utils/", "utils", "UtilsFile"}
		fUtils.Parse()

		fDB := FileParse{"DB open", *newApp, "db/", "db", "DBFile"}
		fDB.Parse()
	}
}
