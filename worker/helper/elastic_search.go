package helper

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var indexName = "client_and_node_decentralization"
var pathDir = "/home/evorun/workspace/decentralized_dashboard/data/transformed"
var elasticHost = "localhost:9200"
var idStart = 14366

type elasticIndex struct {
	Index struct {
		Index string `json:"_index"`
		Id    string `json:"id"`
	} `json:"index"`
}

// FormatBulkUpload adheres to the API specified here:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-bulk.html
func FormatBulkUpload() {
	entries, err := ioutil.ReadDir(pathDir)

	if err != nil {
		fmt.Printf("Failed to read directory %s", err)
	}

	counter := idStart
	for index, entry := range entries {
		absPath := fmt.Sprintf("%v/%v", pathDir, entry)
		indexInfo := elasticIndex{
			Index{
				Index
			}
		}
		//
	}

}

func PutToBulkUpload(fileName string) {
	url := fmt.Sprintf("%v/%v", elasticHost, indexName)
	file, err := os.Open(fileName)
	req, err := http.NewRequest("PUT", url, file)

	if err != nil {
		fmt.Printf("Failed to Make Request. Error: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")
}
