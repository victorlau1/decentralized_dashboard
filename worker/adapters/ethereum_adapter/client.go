package ethereum_adapter

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
    BaseURLV1 = "https://api.etherscan.io/api?module=block&action=getblockreward&blockno=2165405&apikey=U6HY5ZDYB2DRSTP8NPW5FMDIKY2H8QYF4K"
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