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

type MailFileInfo struct {
	nk       string
	devId    string
	clientId string
	time     string
}

// Parse a given directory
func parseDirMail(path string) {

	log.Println("=============================================================")
	log.Println("Processing <mail> directory... :", path)
	log.Println()

	// Check if already processed
	if config.CheckIfAlreadyProccessed && isDirExist(filepath.Join(path, config.DBizedPath)) {
		log.Println("This directory has already been processed. Remove the result directory and retry.")
		log.Println()
		log.Println("Processing <mail> rejected :", path)
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
		result := parseFileMail(filepath.Join(item.Name()))
		if result {
			success++
		} else {
			fail++
		}
	}

	log.Println()
	log.Println("Merging files...")

	// merged, attaches, statistics := mergeFilesMail(config.DBizedPath, 0)

	log.Println()
	log.Println("Summarizing result...")
	// summarizeMailProcess(
	// 	filepath.Join(path, config.DBizedPath, config.SummaryFileName),
	// 	statistics[0],
	// 	statistics[1],
	// 	statistics[2],
	// 	statistics[3],
	// 	attaches,
	// 	fail,
	// )

	log.Println()
	log.Println("Processing <mail> completed :", path)
	log.Printf("Parsed %v file(s), Failed %v file(s)", success, fail)
	log.Println("=============================================================")
}

// func summarizeMailProcess(path string, nks, devs, clients, files, attaches, fail int) {

// 	file, openErr := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
// 	if openErr != nil {
// 		log.Print(openErr)
// 		return
// 	}

// 	defer func() {
// 		if closeErr := file.Close(); closeErr != nil {
// 			log.Println(closeErr)
// 		}
// 	}()

// 	file.WriteString(fmt.Sprintf(
// 		"\n-- Mail log summary --\n\nTotal users (netkey): %v\nTotal mail accounts: %v\nTotal clients: %v\nTotal attachments: %v\n\nTotal log files: %v\nFailed to process: %v\n",
// 		nks, devs, clients, attaches, files, fail))
// }

// Parse a given file
func parseFileMail(name string) bool {
	log.Println("parsing:", name)

	info, nameErr := extractFileNameMail(name)
	if nameErr != nil {
		log.Print(nameErr)
		return false
	}

	destDirPath := path.Join(config.DBizedPath, info.nk, info.devId, strings.ReplaceAll(info.clientId, "...", "etc"))
	dirErr := os.MkdirAll(destDirPath, os.ModePerm)
	if dirErr != nil {
		log.Print(dirErr)
		return false
	}

	moveErr := os.Rename(name, path.Join(destDirPath, info.time+".eml"))
	if moveErr != nil {
		log.Print(moveErr)
	}

	return true
}

func extractFileNameMail(name string) (*MailFileInfo, error) {

	name = strings.ReplaceAll(name, config.LogFileNameSplitter+config.LogFileNameSplitter, config.LogFileNameSplitter+config.PseudoName+config.LogFileNameSplitter)

	pieces := strings.Split(name, config.LogFileNameSplitter)
	if len(pieces) != 6 {
		return nil, fmt.Errorf("invalid mail file name")
	}

	pieces[5] = pieces[5][:len(pieces[5])-4]

	return &MailFileInfo{pieces[0], pieces[3], pieces[4], pieces[5]}, nil
}
