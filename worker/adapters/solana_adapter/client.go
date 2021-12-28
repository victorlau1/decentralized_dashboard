package solanaadapter

import (
	"context"
	"fmt"
	"time"
	"github.com/victorlau1/worker/models"
	// "reflect"
	"github.com/spf13/viper"
	sw "github.com/victorlau1/solanaclient"
	"os"
	"encoding/json"
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

func NewClient() solanaClient {
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

func (s solanaClient) GetOwnershipDecentralization() []models.OwnershipDecentralization {
	auth := s.GetAuthorizationHeaders()
	request := s.client.AccountApi.FetchAccounts(auth)
	Time := time.Now()
	encoding, _ := Time.MarshalJSON()
	encoding[0] = 95
	encoding[14] = 95
	encoding[17] = 95
	encoding[31] = 95
	encoding[34] = 95
	resp, _, _ := request.Limit(1000).Execute()
	for i := 0; i < 1000; i++ {
		var r models.OwnershipDecentralization
		r.Address = *resp[i].Pubkey.Address
		r.Balance = *resp[i].Balance
		r.Blockchain = "Solana"
		r.TimeStamp = Time
		res, _ := json.Marshal(r)
		fmt.Println(string(res))
		output := fmt.Sprintf("data/ownership_decentralization/solana/%s%s.json", r.Address, string(encoding))
		// fmt.Println(output)
		os.WriteFile(output, res, 0644)
	}
	
	return []models.OwnershipDecentralization{models.OwnershipDecentralization{}}
}