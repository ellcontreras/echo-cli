package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"log"
	"strings"
	"os"
	"echo-cli/strs"
	"echo-cli/utils"
)

type FileParse struct {
	name		string
	appName		string
	subFolder	string
	fileName	string
	strName		string
}

func (f *FileParse) Parse(){
	path := f.subFolder+f.fileName + ".go"
	log.Println(path)
	strParse := strings.Replace(strs.Paths[path], "appName", f.appName,-1)
	d1 := []byte(strParse)
	os.MkdirAll(f.appName, os.ModePerm)

	if utils.CheckNil(f.subFolder) {
		os.MkdirAll(f.appName + "/" + f.subFolder, os.ModePerm)
		err := ioutil.WriteFile(f.appName+"/" + f.subFolder + f.fileName + ".go", d1, 0644)
		utils.CheckErr(err)
	} else {
		err := ioutil.WriteFile(f.appName+"/" + f.fileName + ".go", d1, 0644)
		utils.CheckErr(err)
	}
}

func main() {
	fmt.Println("Echo Framework CLI")
	fmt.Println("\n\tUsage: -<command>=<value>")
	fmt.Println("\tExample: -new=polls\n ")

	newApp := flag.String("new", "nameApp", "Generate de folder structure of an app")
	//generate := flag.String("generate", "", "Generate something such as controller, model, route")
	flag.PrintDefaults()
	flag.Parse()

	if utils.CheckNull(newApp) {
		fMain := FileParse{"Main File", *newApp, "", "main", "MainFile"}
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
