package bitcoin_adapter

import (
	// "context"
	// "fmt"
	// "context"
	// "net/url"
	"net/http"
	// "bytes"
	// "io"
	// "errors"
	// "encoding/json"
	// "github.com/victorlau1/worker/models"
	"time"
	// "github.com/spf13/viper"
)

const (
    BaseURLV1 = "https://api.blockchair.com/bitcoin/addresses?offset=100&limit=100"
)

type Client struct {
    BaseURL    string
    apiKey     string
    HTTPClient *http.Client
}

func NewClient(apiKey string) *Client {
    return &Client{
        BaseURL: BaseURLV1,
        apiKey:  apiKey,
        HTTPClient: &http.Client{
            Timeout: time.Minute,
        },
    }
}

