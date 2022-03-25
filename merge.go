package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func mergeFiles(path string, depth int) (int, int, []int) {

	statistics := make([]int, 5)

	items, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println(err)
		return 0, 0, statistics
	}

	statistics[depth] = len(items)

	sort.Slice(items, func(i, j int) bool {
		return items[i].Name() < items[j].Name()
	})

	attachments := []string{}
	mergedText := ""

	var merged int
	var attaches int
	for _, item := range items {
		fullpath := filepath.Join(path, item.Name())
		if item.IsDir() {
			_merged, _attaches, _statistics := mergeFiles(fullpath, depth+1)
			merged += _merged
			attaches += _attaches
			for i, stat := range _statistics {
				statistics[i] += stat
			}
		} else {

			// if attachment
			if len(item.Name()) != 18 {
				attachmentName := item.Name()[16:]
				os.Rename(fullpath, filepath.Join(path, attachmentName))
				attachments = append(attachments, attachmentName)
				continue
			}

			// if chat
			text, readErr := readTextFile(fullpath)
			if readErr != nil {
				log.Print(readErr)
				continue
			}

			if mergedText == "" {
				log.Println("merging:", path)
			}

			mergedText += text

			removeErr := os.Remove(fullpath)
			if removeErr != nil {
				log.Print(removeErr)
			}

			merged++
		}
	}

	if mergedText != "" && len(attachments) > 0 {
		for _, attachment := range attachments {
			mergedText = strings.Replace(mergedText, attachment, config.AttachmentSign+attachment, 1)
		}
	}

	if mergedText != "" {
		mergePath := filepath.Join(path, config.MergeFileName)
		err := writeTextFile(mergePath, mergedText)
		if err != nil {
			log.Print(err)
		}

		indexSkypeFile(&ESDocument{filepath.Join(currentDatePath, mergePath), mergedText})
	}

	return merged, attaches + len(attachments), statistics
}
