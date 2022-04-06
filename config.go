package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	BasePath                 string `json:"BasePath"`
	DBizedPath               string `json:"DBizedPath"`
	CheckIfAlreadyProccessed bool   `json:"CheckIfAlreadyProccessed"`
	MergeFileName            string `json:"MergeFileName"`
	SummaryFileName          string `json:"SummaryFileName"`
	LogFileNameSplitter      string `json:"LogFileNameSplitter"`
	AttachmentPrefix         string `json:"AttachmentPrefix"`
	PseudoName               string `json:"PseudoName"`
	DevShortName             string `json:"DevShortName"`
	ClientShortName          string `json:"ClientShortName"`
	AttachmentSign           string `json:"AttachmentSign"`
	ESIndexSkype             string `json:"ESIndexSkype"`
	ESIndexMail              string `json:"ESIndexMail"`
	EnableElasticSearch      bool   `json:"EnableElasticSearch"`
}

var config Config

func parseConfigFile(configFilePath string) {
	log.Println("Getting configs...")
	jsonFile, err := os.Open(configFilePath)

	if err != nil {
		log.Println("Error while opening config file")
		os.Exit(2)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &config)
}
