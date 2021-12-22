package elastic

import (
    "fmt"
    // "io/ioutil"
    "log"
    // "net/http"
    "context"
    // "os"
    "strconv"
    "github.com/elastic/go-elasticsearch/v8"
    "github.com/elastic/go-elasticsearch/v8/esapi"
    "github.com/victorlau1/worker/models"
    "strings"
    "reflect"
    "encoding/json"
    // "time"
)

func JsonStruct(doc *models.BlockDecentralization) string {

    // Create struct instance of the Elasticsearch fields struct object

    docStruct := &models.BlockDecentralization{
        BlockNumber: doc.BlockNumber,
        TimeStamp: doc.TimeStamp,
        BlockMiner: doc.BlockMiner,
        Blockchain: doc.Blockchain,
    }

    fmt.Println("\ndocStruct:", docStruct)
    fmt.Println("docStruct TYPE:", reflect.TypeOf(docStruct))

    // Marshal the struct to JSON and check for errors
    b, err := json.Marshal(docStruct)
    if err != nil {
        fmt.Println("json.Marshal ERROR:", err)
        return string(err.Error())
    }
    // fmt.Println(string(b))
    return string(b)
}

func Writer(stru *models.BlockDecentralization){
    ctx := context.Background()
    cfg := elasticsearch.Config{
        Addresses: []string{
            "http://localhost:9200",
        },
        // Username: "user",
        // Password: "pass",
    }
    client, err := elasticsearch.NewClient(cfg)
    if err != nil {
        fmt.Println("Elasticsearch connection error:", err)
    }
    bod := JsonStruct(stru)
    fmt.Println(bod)
    req := esapi.IndexRequest{
        Index:      "block_decentralization",
        DocumentID: strconv.Itoa(-1 * stru.BlockNumber - 1),
        Body:       strings.NewReader(bod),
        Refresh:    "true",
    }
    res, err := req.Do(ctx, client)
    if err != nil {
        log.Fatalf("IndexRequest ERROR: %s", err)
    }
    defer res.Body.Close()
    if res.IsError() {
        log.Printf("%s ERROR indexing document ID=%d", res.Status(), stru.BlockNumber)
    } else {

        // Deserialize the response into a map.
        var resMap map[string]interface{}
        if err := json.NewDecoder(res.Body).Decode(&resMap); err != nil {
            log.Printf("Error parsing the response body: %s", err)
        } else {
            log.Printf("\nIndexRequest() RESPONSE:")
            // Print the response status and indexed document version.
            fmt.Println("Status:", res.Status())
            fmt.Println("Result:", resMap["result"])
            fmt.Println("Version:", int(resMap["_version"].(float64)))
            fmt.Println("resMap:", resMap)
            // fmt.Println("\n")
        }
    }
}