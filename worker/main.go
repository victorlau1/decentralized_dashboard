package main

import (
	"fmt"
	sol "github.com/victorlau1/worker/adapters/solana_adapter"
	// bit "github.com/victorlau1/worker/adapters/bitcoin_adapter"
	// eth "github.com/victorlau1/worker/adapters/ethereum_adapter"
	// git "github.com/victorlau1/worker/adapters/github_adapter"
	"github.com/victorlau1/worker/models"
	// "reflect"
    "encoding/json"
	"io/ioutil"
	"log"
	// "time"
	elastic "github.com/victorlau1/worker/elastic"
	"github.com/spf13/viper"
	// "os"
	// "strconv"
	// "path/filepath"
	// "context"
)

func main() {
	WorkerConfig()
	
	



	c := sol.NewClient()
	// ctx := context.Background()
	c.GetOwnershipDecentralization()

	// currentDirectory, err := os.Getwd()
    // if err != nil {
    //     log.Fatal(err)
    // }
	// fmt.Println(currentDirectory)
    // iterate("data/github/erigon")

	// c := bit.NewClient(viper.GetString("ETHEREUM_API_KEY"))
    // ctx := context.Background()
	// c.GetDevDecentralization(ctx)


	// for i := 1; i < 172; i++ {
	// 	content, err := ioutil.ReadFile(fmt.Sprintf("%s%d.json", "raw_data/github/solana/page_", i))
	// 	if err != nil {
	// 		log.Fatal("Error when opening file: ", err)
	// 	}
	// 	var payload git.Commits
	// 	err = json.Unmarshal(content, &payload)
	// 	if err != nil {
	// 		log.Fatal("Error during Unmarshal(): ", err)
	// 	}
	// 	// for j := 0; j < len(payload); j++ {
	// 	for j := 0; j < len(payload); j++ {
	// 		var doc models.DevDecentralization
	// 		doc.CommitID = payload[j].Sha
	// 		doc.Committer = payload[j].Commit.Committer.Name
	// 		doc.Email = payload[j].Commit.Committer.Email
	// 		doc.TimeStamp = payload[j].Commit.Committer.Date
	// 		doc.Client = "SolanaClient"
	// 		doc.Blockchain = "Solana"
	// 		data, _ := json.Marshal(doc)
	// 		fmt.Println(string(data))
	// 		os.WriteFile(fmt.Sprintf("%s%s/%s.json", "data/github/", "solana" , doc.CommitID), data, 0644)
	// 		// elastic.Writer(&doc)
	// 	}
	// }
	

	// fmt.Println(res)
	
	// if err != nil {
	// 	fmt.Println(err)
	// }
	
	// start := time.Now()
	// for i := 715981; i > 713965; i-- {
	// // for i := 715981; i > 715980; i-- {
	// 	fmt.Println(i)
	// 	res, err := c.GetBlockDecentralization(ctx, i)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		break
	// 	}
	// 	// elastic.Writer(res)
	// 	fmt.Println(res)
	// }
	// end := time.Now()
	// fmt.Println(end.Sub(start))

	
}

func iterate(folder_path string) {

	items, _ := ioutil.ReadDir(folder_path)
    for _, item := range items {
		// handle file there
		sendToElastic(fmt.Sprintf("%s/%s", folder_path, item.Name()))
		// fmt.Println(item.Name())
    }
	
}

func sendToElastic(file string){
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	var payload models.DevDecentralization
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	data, _ := json.Marshal(payload)
	fmt.Println(string(data))
	// os.WriteFile(fmt.Sprintf("%s%s/%s.json", "data/github/", "solana" , doc.CommitID), data, 0644)
	elastic.Writer(&payload)
	
}

// WorkerConfig grabs the necessary settings to boot up the worker
// It should also take in commands to configure the blockchain the worker is requesting against
func WorkerConfig() {
	viper.AutomaticEnv()
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.SetConfigType("dotenv")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}
