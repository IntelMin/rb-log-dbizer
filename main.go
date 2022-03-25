package main

import "path/filepath"

var currentDatePath string = ""

func main() {
	parseConfigFile()
	initElasticSearch()

	currentDatePath = produceDatePath(2022, 1, 8)
	parseDirSkype(filepath.Join(config.BasePath, currentDatePath))

	refreshIndex(config.ESIndexSkype)
}
