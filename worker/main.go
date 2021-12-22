package main

import (
	"fmt"
	// sol "github.com/victorlau1/worker/adapters/solana_adapter"
	bit "github.com/victorlau1/worker/adapters/bitcoin_adapter"
	// eth "github.com/victorlau1/worker/adapters/ethereum_adapter"
	// "github.com/victorlau1/worker/models"
	// "reflect"
    // "encoding/json"
	"time"
	// elastic "github.com/victorlau1/worker/elastic"
	"github.com/spf13/viper"
	"context"
)

func main() {
	WorkerConfig()
	// client := sol.NewClient()
	// client.GetClientsDecentralization()
	c := bit.NewClient(viper.GetString("ETHEREUM_API_KEY"))
    ctx := context.Background()
	// res, err := c.GetBlockDecentralization(ctx, 100000)
	// fmt.Println(res)
	// fmt.Println(err)
	start := time.Now()
	for i := 715234; i > 713218; i-- {
		fmt.Println(i)
		res, err := c.GetBlockDecentralization(ctx, i)
		if err != nil {
			fmt.Println(err)
			break
		}
		// elastic.Writer(res)
		fmt.Println(res)
	}
	end := time.Now()
	fmt.Println(end.Sub(start))

	
}

// func jsonStruct(doc *models.BlockDecentralization) string {

//     // Create struct instance of the Elasticsearch fields struct object

//     docStruct := &models.BlockDecentralization{
//         BlockNumber: doc.BlockNumber,
//         TimeStamp: doc.TimeStamp,
//         BlockMiner: doc.BlockMiner,
//         Blockchain: doc.Blockchain,
//     }

//     fmt.Println("\ndocStruct:", docStruct)
//     fmt.Println("docStruct TYPE:", reflect.TypeOf(docStruct))

//     // Marshal the struct to JSON and check for errors
//     b, err := json.Marshal(docStruct)
//     if err != nil {
//         fmt.Println("json.Marshal ERROR:", err)
//         return string(err.Error())
//     }
//     // fmt.Println(string(b))
//     return string(b)
// }

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
