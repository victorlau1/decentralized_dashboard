package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	WorkerConfig()
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
