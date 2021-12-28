package github_adapter

import (
	// "context"
	"fmt"
	// "context"
	// "net/url"
	"net/http"
	// "bytes"
	// "io"
	"errors"
	"encoding/json"
	// "github.com/victorlau1/worker/models"
	"time"
	// "github.com/spf13/viper"
)

const (
    BaseURLV1 = "https://api.github.com/repos/bitcoin/bitcoin/commits?per_page=100&page="
	BaseURLV2 = "https://api.github.com/repos/ethereum/go-ethereum/commits?per_page=100&page="
	BaseURLV3 = "https://api.github.com/repos/openethereum/openethereum/commits?per_page=100&page="
	BaseURLV4 = "https://api.github.com/repos/ledgerwatch/erigon/commits?per_page=100&page="
	BaseURLV5 = "https://api.github.com/repos/NethermindEth/nethermind/commits?per_page=100&page="
	BaseURLV6 = "https://api.github.com/repos/solana-labs/solana/commits?per_page=100&page="
)

type errorResponse struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

type Client struct {
    BaseURL    string
    apiKey     string
    HTTPClient *http.Client
	client	   string
}

func NewClient(apiKey string, client string) *Client {
	var BaseURL string
	switch client {
	case "bitcoin":
		BaseURL = BaseURLV1
	case "geth":
		BaseURL = BaseURLV2
	case "openethereum":
		BaseURL = BaseURLV3
	case "erigon":
		BaseURL = BaseURLV4
	case "nethermind":
		BaseURL = BaseURLV5
	case "solana":
		BaseURL = BaseURLV6
	default:
		BaseURL = BaseURLV1
	}
    return &Client{
        BaseURL: BaseURL,
        apiKey:  apiKey,
        HTTPClient: &http.Client{
            Timeout: time.Minute,
        },
		client: client,
    }
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
    req.Header.Set("Content-Type", "application/json; charset=utf-8")
    req.Header.Set("Accept", "application/json; charset=utf-8")
    // req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

    res, err := c.HTTPClient.Do(req)
    if err != nil {
        return err
    }
	// fmt.Println(res.Body)
    defer res.Body.Close()

    if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
        var errRes errorResponse
        if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
            return errors.New(errRes.Message)
        }

        return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
    }
    if err = json.NewDecoder(res.Body).Decode(v); err != nil {
        return err
    }

    return nil
}
