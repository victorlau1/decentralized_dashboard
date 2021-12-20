package main

import (
	"fmt"
	// sol "github.com/victorlau1/worker/adapters/solana_adapter"
	bit "github.com/victorlau1/worker/adapters/bitcoin_adapter"
	"github.com/spf13/viper"
	"context"
)

func main() {
	WorkerConfig()
	// client := sol.NewClient()
	// client.GetClientsDecentralization()
	c := bit.NewClient("")
    ctx := context.Background()
    res, err := c.GetOwnershipDecentralization(ctx)
	fmt.Println(res)
	fmt.Println(err)
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
