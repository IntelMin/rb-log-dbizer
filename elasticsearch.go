package main

import (
	"context"
	"encoding/json"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v7"
)

type ESDocument struct {
	path string
	text string
}

var es *elasticsearch.Client = nil

func initElasticSearch() {
	log.Println("Initializing Elastic Search...")

	var r map[string]interface{}

	_es, err := elasticsearch.NewDefaultClient()

	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := _es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	// Check response status
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}
	// Deserialize the response into a map.
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print client and server version numbers.
	log.Printf("Client: %s", elasticsearch.Version)
	log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
	log.Println(strings.Repeat("~", 37))

	es = _es
}

func refreshIndex(index string) {

	req := esapi.IndicesRefreshRequest{
		Index: []string{index},
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error refreshing index: %s", res.Status(), index)
	}
}

func indexFile(document *ESDocument, index string) bool {

	re := regexp.MustCompile(`\r?\n`)
	filteredText := re.ReplaceAllString(document.text, " ")
	filteredText = strings.ReplaceAll(filteredText, "\"", "'")
	// Build the request body.
	var b strings.Builder
	b.WriteString(`{"path" : "`)
	b.WriteString(document.path)
	b.WriteString(`",`)
	b.WriteString(`"text" : "`)
	b.WriteString(filteredText)
	b.WriteString(`",`)
	b.WriteString(`"timestamp" : "`)
	b.WriteString(strconv.FormatInt(time.Now().Unix(), 10))
	b.WriteString(`"}`)

	// log.Println("=================================================================================")
	// log.Print(document.text)
	// log.Println("--------------------------------------")
	// log.Print(filteredText)
	// log.Println("--------------------------------------")
	// log.Print(b.String())
	// log.Println("======================================")

	// Set up the request object.
	req := esapi.IndexRequest{
		Index: index,
		Body:  strings.NewReader(b.String()),
	}

	// Perform the request with the client.
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return false
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error indexing document in path: %s", res.Status(), document.path)
		return false
	}

	return true
}
