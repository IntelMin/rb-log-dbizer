package main

import (
	"io/ioutil"
	"log"
	"os"
)

func isDirExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func readTextFile(path string) (string, error) {

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Println(err)
		}
	}()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content[:]), nil
}
