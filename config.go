package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	BasePath          string `json:"BasePath"`
	DBizedPath        string `json:"DBizedPath"`
	MergeFileName     string `json:"MergeFileName"`
	SummaryFileName   string `json:"SummaryFileName"`
	SkypeNameSplitter string `json:"SkypeNameSplitter"`
	PseudoName        string `json:"PseudoName"`
	DevShortName      string `json:"DevShortName"`
	ClientShortName   string `json:"ClientShortName"`
	AttachmentSign    string `json:"AttachmentSign"`
	ESIndexSkype      string `json:"ESIndexSkype"`
}

var config Config

func parseConfigFile() {
	log.Println("Getting configs...")
	jsonFile, err := os.Open("config.json")

	if err != nil {
		log.Println("Error while opening config file")
		return
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &config)

	log.Print(config)
}
