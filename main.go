package main

import "path/filepath"

var currentDatePath string = ""

func main() {
	initElasticSearch()

	currentDatePath = produceDatePath(2022, 1, 8)
	parseDirSkype(filepath.Join(BasePath, currentDatePath))

	refreshIndex(ESIndexSkype)
}
