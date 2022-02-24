package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func mergeFiles(path string, depth int) (int, []int) {

	statistics := make([]int, 5)

	items, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println(err)
		return 0, statistics
	}

	statistics[depth] = len(items)

	sort.Slice(items, func(i, j int) bool {
		return items[i].Name() < items[j].Name()
	})

	var merged int
	var mergeFile *os.File
	for _, item := range items {
		if item.IsDir() {
			_merged, _statistics := mergeFiles(filepath.Join(path, item.Name()), depth+1)
			merged += _merged
			for i, stat := range _statistics {
				statistics[i] += stat
			}
		} else {
			text, readErr := readTextFile(filepath.Join(path, item.Name()))
			if readErr != nil {
				log.Print(readErr)
				continue
			}

			if mergeFile == nil {
				log.Println("merging:", path)

				file, openErr := os.OpenFile(filepath.Join(path, MergeFileName), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
				if openErr != nil {
					log.Print(openErr)
					continue
				}
				mergeFile = file

				defer func() {
					if closeErr := mergeFile.Close(); closeErr != nil {
						log.Println(closeErr)
					}
				}()
			}

			_, writeErr := mergeFile.WriteString(text)
			if writeErr != nil {
				log.Print(writeErr)
				continue
			}
			merged++
		}
	}

	return merged, statistics
}
