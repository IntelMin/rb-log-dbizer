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

func readBytesFile(path string) ([]byte, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Println(err)
		}
	}()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func writeTextFile(path string, text string) error {
	file, openErr := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if openErr != nil {
		return openErr
	}

	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			log.Println(closeErr)
		}
	}()

	_, writeErr := file.WriteString(text)
	if writeErr != nil {
		return writeErr
	}

	return nil
}
