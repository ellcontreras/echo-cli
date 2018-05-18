package main

import (
	"github.com/Obsinqsob01/echo-cli/strs"
	"github.com/Obsinqsob01/echo-cli/utils"
	"flag"
	"io/ioutil"
	"os"
	"strings"
	"fmt"
	"io"
	"bufio"
)

type FileParse struct {
	name		string
	appName		string
	subFolder	string
	fileName	string
	strName		string
}

func (f *FileParse) Parse() {
	path := f.subFolder+f.fileName + ".go"

	fmt.Println("\tCreando...", path)

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
	stream := io.Reader(os.Stdin)
	reader := bufio.NewReader(stream)

	newApp := flag.String("new", "", "Generate de folder structure of an app")
	generate := flag.String("generate", "", "Generate something such as action, controller, model, route")

	flag.Parse()

	fmt.Println(*newApp, *generate)

	if utils.CheckNil(*newApp) {
		fmt.Println("Echo Framework CLI")
		fmt.Println("\n\tUsage: -<command>=<value>")
		fmt.Println("\tExample: -new=polls\n ")

		flag.PrintDefaults()
	} else {
		fmt.Println("Generando proyecto", *newApp)
	}

	if !utils.CheckNil(*newApp) {
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

	if !utils.CheckNil(*generate) {
		switch *generate {
		case "action":
			fmt.Println("Escribe el nombre del action: ")
			name, err := reader.ReadString('\n')
			utils.CheckErr(err)

			af := FileParse{"Action", name, "actions/", name, "Action"}
			af.Parse()
			break
		}
	}
}
