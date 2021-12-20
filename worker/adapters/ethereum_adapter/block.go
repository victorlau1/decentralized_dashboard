package ethereum_adapter

import (
	// "context"
	"fmt"
	"context"
	// "net/url"
	"net/http"
	// "bytes"
	// "io"
	"errors"
	"encoding/json"
	// "github.com/victorlau1/worker/models"
	// "github.com/spf13/viper"
)

type errorResponse struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

type BlockStats struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  struct {
		BlockNumber          string        `json:"blockNumber"`
		TimeStamp            string        `json:"timeStamp"`
		BlockMiner           string        `json:"blockMiner"`
		BlockReward          string        `json:"blockReward"`
		Uncles               []interface{} `json:"uncles"`
		UncleInclusionReward string        `json:"uncleInclusionReward"`
	} `json:"result"`
}


func (c *Client) GetBlock(ctx context.Context) (*BlockStats, error) {
    // limit := 100
    // page := 1
    // if options != nil {
    //     limit = options.Limit
    //     page = options.Page
    // }

    req, err := http.NewRequest("GET", c.BaseURL, nil)
    if err != nil {
        return nil, err
    }

    req = req.WithContext(ctx)

    res := BlockStats{}
    if err := c.sendRequest(req, &res); err != nil {
        return nil, err
    }

    return &res, nil
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
    req.Header.Set("Content-Type", "application/json; charset=utf-8")
    req.Header.Set("Accept", "application/json; charset=utf-8")
    // req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

    res, err := c.HTTPClient.Do(req)
    if err != nil {
        return err
    }
	fmt.Println(res.Body)
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