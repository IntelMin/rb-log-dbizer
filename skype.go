package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type SkypeFileInfo struct {
	nk       string
	devId    string
	clientId string
	time     string
	filename string
}

// Parse a given directory
func parseDirSkype(path string) {

	log.Println("=============================================================")
	log.Println("Processing <skype> directory... :", path)
	log.Println()

	// Check if already processed
	if isDirExist(filepath.Join(path, config.DBizedPath)) {
		log.Println("This directory has already been processed. Remove the result directory and retry.")
		log.Println()
		log.Println("Processing <skype> rejected :", path)
		log.Printf("Parsed %v file(s), Failed %v file(s), Merged %v file(s)", 0, 0, 0)
		log.Println("=============================================================")
		return
	}

	// Reading directory
	items, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println(err)
		return
	}

	// Changing working directory
	pathErr := os.Chdir(path)
	if pathErr != nil {
		log.Println(pathErr)
		return
	}

	log.Println("Parsing files...")

	var success, fail int
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		result := parseFileSkype(filepath.Join(item.Name()))
		if result {
			success++
		} else {
			fail++
		}
	}

	log.Println()
	log.Println("Merging files...")

	merged, attaches, statistics := mergeFiles(config.DBizedPath, 0)

	log.Println()
	log.Println("Summarizing result...")
	summarizeSkypeProcess(
		filepath.Join(path, config.DBizedPath, config.SummaryFileName),
		statistics[0],
		statistics[1],
		statistics[2],
		statistics[3],
		attaches,
		fail,
	)

	log.Println()
	log.Println("Processing <skype> completed :", path)
	log.Printf("Parsed %v file(s), Failed %v file(s), Merged %v file(s)", success, fail, merged)
	log.Println("=============================================================")
}

func summarizeSkypeProcess(path string, nks, devs, clients, files, attaches, fail int) {

	file, openErr := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if openErr != nil {
		log.Print(openErr)
		return
	}

	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			log.Println(closeErr)
		}
	}()

	file.WriteString(fmt.Sprintf(
		"\n-- Skype log summary --\n\nTotal users (netkey): %v\nTotal skype accounts: %v\nTotal clients: %v\nTotal attachments: %v\n\nTotal log files: %v\nFailed to process: %v\n",
		nks, devs, clients, attaches, files, fail))
}

// Parse a given file
func parseFileSkype(name string) bool {
	log.Println("parsing:", name)

	info, nameErr := extractFileNameSkype(name)
	if nameErr != nil {
		log.Print(nameErr)
		return false
	}

	destDirPath := path.Join(config.DBizedPath, info.nk, info.devId, info.clientId)
	dirErr := os.MkdirAll(destDirPath, os.ModePerm)
	if dirErr != nil {
		log.Print(dirErr)
		return false
	}

	if info.filename != "" {
		// Moving attachment
		moveError := os.Rename(name, path.Join(destDirPath, info.time+"__"+info.filename))
		if moveError != nil {
			log.Print(moveError)
			return false
		}
	} else {
		text, readErr := readTextFile(name)
		if readErr != nil {
			log.Print(readErr)
			return false
		}

		// Moving txt file
		SNS := " "
		text = strings.ReplaceAll(text, SNS+SNS+":", SNS+config.PseudoName+SNS+":")
		text = strings.ReplaceAll(text, SNS+info.devId+SNS+":", SNS+config.DevShortName+SNS+":")
		text = strings.ReplaceAll(text, SNS+info.clientId+SNS+":", SNS+config.ClientShortName+SNS+":")

		writeErr := writeTextFile(path.Join(destDirPath, info.time+".txt"), text)
		if writeErr != nil {
			log.Print(writeErr)
		}

		removeErr := os.Remove(name)
		if removeErr != nil {
			log.Print(removeErr)
		}
	}

	return true
}

func extractFileNameSkype(name string) (*SkypeFileInfo, error) {

	name = strings.ReplaceAll(name, config.SkypeNameSplitter+config.SkypeNameSplitter, config.SkypeNameSplitter+config.PseudoName+config.SkypeNameSplitter)

	pieces := strings.Split(name, config.SkypeNameSplitter)
	if len(pieces) != 6 {
		return nil, fmt.Errorf("Invalid skype file name!")
	}

	filename := ""
	// Check if file
	if len(pieces[5]) > 18 { // time + .txt = 18 length
		filename = pieces[5][16:]
		pieces[5] = pieces[5][:14]
	} else {
		// Remove .txt (extension) from time string
		pieces[5] = pieces[5][:len(pieces[5])-4]
	}
	return &SkypeFileInfo{pieces[0], pieces[3], pieces[4], pieces[5], filename}, nil
}
