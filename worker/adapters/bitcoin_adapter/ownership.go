package bitcoin_adapter

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

type Owner struct {
	Data []struct {
		Address string `json:"address"`
		Balance int64  `json:"balance"`
	} `json:"data"`
	Context struct {
		Code           int    `json:"code"`
		Source         string `json:"source"`
		Limit          int    `json:"limit"`
		Offset         int    `json:"offset"`
		Rows           int    `json:"rows"`
		PreRows        int    `json:"pre_rows"`
		TotalRows      int    `json:"total_rows"`
		State          int    `json:"state"`
		MarketPriceUsd float64    `json:"market_price_usd"`
		Cache          struct {
			Live     bool        `json:"live"`
			Duration int         `json:"duration"`
			Since    string      `json:"since"`
			Until    string      `json:"until"`
			Time     float64 `json:"time"`
		} `json:"cache"`
		API struct {
			Version         string      `json:"version"`
			LastMajorUpdate string      `json:"last_major_update"`
			NextMajorUpdate interface{} `json:"next_major_update"`
			Documentation   string      `json:"documentation"`
			Notice          string      `json:"notice"`
		} `json:"api"`
		Server      string  `json:"server"`
		Time        float64 `json:"time"`
		RenderTime  float64 `json:"render_time"`
		FullTime    float64 `json:"full_time"`
		RequestCost int     `json:"request_cost"`
	} `json:"context"`
}

func (c *Client) GetOwnershipDecentralization(ctx context.Context) (*Owner, error) {
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

    res := Owner{}
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