package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Obsinqsob01/echo-cli/strs"
	"github.com/Obsinqsob01/echo-cli/utils"
)

type FileParse struct {
	name      string
	appName   string
	subFolder string
	fileName  string
	strName   string
}

//Parse parse the files
func (f *FileParse) Parse() {
	path := f.subFolder + f.fileName + ".go"

	strParse := strings.Replace(strs.Paths[path], "appName", f.appName, -1)
	d1 := []byte(strParse)
	os.MkdirAll(f.appName, os.ModePerm)

	if utils.CheckNil(f.subFolder) {
		fmt.Println("\tCreating...", path)
		os.MkdirAll(f.appName+"/"+f.subFolder, os.ModePerm)
		err := ioutil.WriteFile(f.appName+"/"+f.subFolder+f.fileName+".go", d1, 0644)
		utils.CheckErr(err)
	} else {
		err := ioutil.WriteFile(f.appName+"/"+f.fileName+".go", d1, 0644)
		utils.CheckErr(err)
	}
}

func main() {
	//Initialize the Path map to get the string data
	strs.Init()

	stream := io.Reader(os.Stdin)
	reader := bufio.NewReader(stream)

	newApp := flag.String("new", "", "Generate de folder structure of an app")
	generate := flag.String("generate", "", "Generate something such as action, controller, model, route")

	flag.Parse()

	if !utils.CheckNil(*newApp) && !utils.CheckNil(*generate) {
		fmt.Println("Echo Framework CLI")
		fmt.Println("\n\tUsage: -<command>=<value>")
		fmt.Println("\tExample: -new polls\n ")

		flag.PrintDefaults()
	}

	if utils.CheckNil(*newApp) {
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

	//Check if the generate flag is null or not
	if utils.CheckNil(*generate) {
		switch *generate {
		case "action":
			fmt.Println("Write the action's name: ")
			name, err := reader.ReadString('\n')
			utils.CheckErr(err)

			af := FileParse{"Action", name, "actions/", name, "Action"}
			af.Parse()
			break
		}
	}
}
