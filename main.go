package main

import (
	"log"
	"path/filepath"
)

var currentDatePath string = ""

var targets = []string{"skype", "mail"}

func main() {

	args := parseArgs()

	parseConfigFile(args.configPath)

	if config.EnableElasticSearch {
		initElasticSearch()
	} else {
		log.Println("Skipping elastic search configuration...")
	}

	currentDatePath = args.dateStr

	if args.target == "skype" {
		parseDirSkype(filepath.Join(config.BasePath, "skype", currentDatePath))
	} else if args.target == "mail" {
		parseDirMail(filepath.Join(config.BasePath, "mail", currentDatePath))
	}

	if config.EnableElasticSearch {
		refreshIndex(config.ESIndexSkype)
	}
}
