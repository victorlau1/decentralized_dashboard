package solanaadapter

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/victorlau1/worker/models"

	"github.com/spf13/viper"
	sw "github.com/victorlau1/solanaclient"
)

type solanaClient struct {
	client          *sw.APIClient
	transformations Transformations
}

type solanaConfiguration struct {
	configuration *sw.Configuration
}

// SolanaClient is the client used to retrieve the data from the SolanaBlockchain
// The client must return the models in the expected structure for ElasticSearch consumption
type SolanaClient interface {
	GetClientsDecentralization() ([]models.ClientDecentralization, error)
	//GetNodeDecentralization() []models.NodeDecentralization
	//GetPOWDecentralization() []models.
	//GetAccountWealthDecentralization()
}

func NewClient(solClient *solanaClient, transformations Transformations) solanaClient {
	if solClient == nil {
		apiClient := sw.NewAPIClient(sw.NewConfiguration())
		return solanaClient{
			client:          apiClient,
			transformations: transformations,
		}
	}
	return (*solClient)
}

func (s solanaClient) GetAuthorizationHeaders() context.Context {
	key := viper.GetString("SOLANA_API_KEY")
	// fmt.Println("key", ":", key)
	return context.WithValue(context.Background(), sw.ContextAccessToken, key)
}

//TODO: Optimize into golang channels using RxGo to improve performance.
//TODO: Handle errors
func (s solanaClient) GetClientsDecentralization() ([]models.ClientDecentralization, error) {
	auth := s.GetAuthorizationHeaders()
	request := s.client.OtherApi.FetchNonValidators(auth)
	resp, _, _ := request.Execute()
	data, err := json.Marshal(resp)

	if err != nil {
		fmt.Println(err)
	}
	current_timestamp := time.Now().Format(time.RFC3339)
	filename := fmt.Sprintf("/home/evorun/workspace/decentralized_dashboard/data/solana_node_data_%v.json", current_timestamp)
	os.WriteFile(filename, data, 0644)
	// schemaModels, err := s.transformations.HandleClientTransformation(resp)

	return []models.ClientDecentralization{}, nil
}

//TODO: Clean this up
func (s solanaClient) GetNodeDecentralization() []models.NodeDecentralization {
	// auth := s.GetAuthorizationHeaders()
	// request := s.client.OtherApi.Fetch
	return []models.NodeDecentralization{}
}

func (s solanaClient) GetValidatorNode(pubkey string) sw.Validator {
	auth := s.GetAuthorizationHeaders()
	request := s.client.ValidatorApi.FetchValidatorByVotepubkey(auth, pubkey)
	resp, _, _ := request.Execute()
	return resp
}
