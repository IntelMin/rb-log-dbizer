package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
)

type Args struct {
	configPath string
	dateStr    string
	target     string
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func parseArgs() Args {
	var configPath string
	var dateStr string
	var target string

	// flags declaration using flag package
	flag.StringVar(&configPath, "c", "", "Specify config file path.")
	flag.StringVar(&dateStr, "d", "", "Specify date string to process. Format: YYYY/MM/DD")
	flag.StringVar(&target, "t", "", "Specify the target service to process. Available: skype|mail")

	flag.Parse() // after declaring flags we need to call it

	if configPath == "" {
		fmt.Println("Config file is not specified. Usage: -c <path/to/config/file>")
		os.Exit(2)
	}

	if dateStr == "" {
		fmt.Println("Date is not specified. Usage: -d <YYYY/MM/DD>")
		os.Exit(2)
	}

	if target == "" {
		fmt.Println("Target is not specified. Usage: -t <skype|mail>")
		os.Exit(2)
	}

	re := regexp.MustCompile(`((19|20)\d\d)\/(0?[1-9]|1[012])\/(0?[1-9]|[12][0-9]|3[01])`)

	if !re.MatchString(dateStr) {
		fmt.Println("Invalid date format. Usage: -d <YYYY/MM/DD>")
		os.Exit(2)
	}

	if !contains(targets, target) {
		fmt.Println("Invalid target specified. Usage: -t <skype|mail>")
		os.Exit(2)
	}

	return Args{configPath, dateStr, target}
}
