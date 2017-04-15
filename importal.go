package main

import (
	"flag"
	"fmt"
	"github.com/petersky/importal/server"
	"os"
	"path"
)

var (
	version     string
	programName = "Importal"
)

func displayVersion() {
	fmt.Println(programName + " Version: " + version)
}

func main() {
	importalHome := os.Getenv("IMPORTAL_HOME")
	webPort := os.Getenv("PORT")
	webIP := os.Getenv("IP")

	if importalHome == "" {
		importalHome = path.Join(os.Getenv("HOME"), ".importal")
	}

	showVersion := flag.Bool("version", false, "Show version")

	fmt.Println("Importal home: " + importalHome)

	flag.Parse()

	if *showVersion {
		displayVersion()
		os.Exit(0)
	}

	server.StartServer(webIP, webPort)
}
