package solanaadapter

import (
	"context"
	"fmt"

	"github.com/victorlau1/worker/models"

	"github.com/spf13/viper"
	sw "github.com/victorlau1/solanaclient"
)

type solanaClient struct {
	client *sw.APIClient
}

type solanaConfiguration struct {
	configuration *sw.Configuration
}

type SolanaClient interface {
	GetClientsDecentralization() []models.ClientDecentralization
	//GetNodeDecentralization() []models.NodeDecentralization
	//GetPOWDecentralization() []models.
	//GetAccountWealthDecentralization()
}

func NewClient() SolanaClient {
	apiClient := sw.NewAPIClient(sw.NewConfiguration())
	return solanaClient{apiClient}
}

func (s solanaClient) GetAuthorizationHeaders() context.Context {
	key := viper.GetString("SOLANA_API_KEY")
	fmt.Println("key", ":", key)
	return context.WithValue(context.Background(), sw.ContextAccessToken, key)
}

func (s solanaClient) GetClientsDecentralization() []models.ClientDecentralization {
	auth := s.GetAuthorizationHeaders()
	request := s.client.OtherApi.FetchNonValidators(auth)
	resp, _, _ := request.Execute()
	fmt.Printf("%+v\n", resp)
	return []models.ClientDecentralization{models.ClientDecentralization{}}
}
