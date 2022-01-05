package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	// "time"

	"github.com/victorlau1/worker/models"
)

var indexName = "dev_decentralization_authors"
var pathDir = "data/transformed"
var outputPath = "data/elastic_bulk_transformed"
var elasticHost = "http://localhost:9200"
var idStart = 1

type elasticIndex struct {
	Index IndexAction `json:"index"`
}

type IndexAction struct {
	Index string `json:"_index"`
	Id    string `json:"_id"`
}

// FormatBulkUpload adheres to the API specified here:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-bulk.html
func FormatBulkUpload() {
	entries, err := ioutil.ReadDir(pathDir)

	if err != nil {
		fmt.Printf("Failed to read directory %s", err)
	}

	for _, entry := range entries {
		b := new(bytes.Buffer)
		enc := json.NewEncoder(b)
		absPath := fmt.Sprintf("%s/%s", pathDir, entry.Name())

		file, err := os.ReadFile(absPath)

		if err != nil {
			fmt.Printf("Failed to read file: %s", err)
		}

		rows := strings.Split(string(file), "\n")
		for _, row := range rows {
			
			nm := &models.DevDecentralization{}
			err = json.Unmarshal([]byte(row), nm)
			indexInfo := elasticIndex{
				Index: IndexAction{
					Index: indexName,
					Id:    fmt.Sprintf("%v_%s", nm.CommitID, nm.Client),
				},
			}
			enc.Encode(indexInfo)
			enc.Encode(nm)
		}
		nf := fmt.Sprintf("%v/%v", outputPath, entry.Name())
		nb, _ := ioutil.ReadAll(b)
		err = os.WriteFile(nf, nb, 0644)

		if err != nil {
			fmt.Printf("Failed to write file: %s", err)
		}
	}
}

func PutToBulkUpload() {
	url := fmt.Sprintf("%v/%v/%v", elasticHost, indexName, "_bulk?pretty")

	entries, err := ioutil.ReadDir(pathDir)

	if err != nil {
		fmt.Printf("Failed to read directory %s", err)
	}

	for _, entry := range entries {
		fileName := fmt.Sprintf("%v/%v", outputPath, entry.Name())
		file, err := os.Open(fileName)

		if err != nil {
			fmt.Printf("Failed to read file")
		}

		client := &http.Client{}

		req, err := http.NewRequest(http.MethodPut, url, file)

		if err != nil {
			fmt.Printf("Failed to Make Request. Error: %v", err)
		}

		req.Header.Add("Content-Type", "application/json")
		// req.Header.Add("Content-Disposition", fmt.Sprintf("attachment; filename=%v", fileName))
		resp, err := client.Do(req)

		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Printf("%s", body)
	}
}