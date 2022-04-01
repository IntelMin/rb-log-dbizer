package main

import (
	"log"
	"path/filepath"
)

var currentDatePath string = ""

func main() {
	parseConfigFile()

	if config.EnableElasticSearch == true {
		initElasticSearch()
	} else {
		log.Println("Skipping elastic search configuration...")
	}

	currentDatePath = produceDatePath(2022, 1, 1)
	parseDirSkype(filepath.Join(config.BasePath, currentDatePath))

	if config.EnableElasticSearch == true {
		refreshIndex(config.ESIndexSkype)
	}
}
