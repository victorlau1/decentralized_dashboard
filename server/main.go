package main

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	sw "github.com/victorlau1/solanaclient"
)

func main() {
	ServerConfig()
	getClients()
}

func ServerConfig() {
	viper.AutomaticEnv()
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.SetConfigType("dotenv")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}

func NewClientWithoutAuth() *sw.APIClient {
	return sw.NewAPIClient(sw.NewConfiguration())
}

func NewAuthorizationHeaders() context.Context {
	key := viper.GetString("SOLANA_API_KEY")
	fmt.Println("key", ":", key)
	return context.WithValue(context.Background(), sw.ContextAccessToken, key)
}

func getClients() {
	client := NewClientWithoutAuth()
	auth := NewAuthorizationHeaders()
	request := client.OtherApi.FetchNonValidators(auth)
	resp, _, _ := request.Execute()
	fmt.Print(resp)
}
