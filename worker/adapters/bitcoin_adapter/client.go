package bitcoin_adapter

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
    BaseURLV1 = "https://api.blockchair.com/bitcoin/addresses?offset=100&limit=100"
	BaseURLV2 = "https://blockchain.info/block-height"
)

type errorResponse struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

type Client struct {
    BaseURL    string
    apiKey     string
    HTTPClient *http.Client
}

func NewClient(apiKey string) *Client {
    return &Client{
        BaseURL: BaseURLV2,
        apiKey:  apiKey,
        HTTPClient: &http.Client{
            Timeout: time.Minute,
        },
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
