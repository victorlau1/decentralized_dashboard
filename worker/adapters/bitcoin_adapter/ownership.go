package bitcoin_adapter

import (
	// "context"
	"fmt"
	"context"
	// "net/url"
	"time"
	"net/http"
	// "bytes"
	// "io"
	// "errors"
	"encoding/json"
	"os"
	"github.com/victorlau1/worker/models"
	// "reflect"
	// "github.com/spf13/viper"
)



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
		RequestCost float64     `json:"request_cost"`
	} `json:"context"`
}

func (c *Client) GetOwnershipDecentralization(ctx context.Context) {
	Time := time.Now()
	encoding, _ := Time.MarshalJSON()
	encoding[0] = 95
	encoding[14] = 95
	encoding[17] = 95
	encoding[31] = 95
	encoding[34] = 95
	fmt.Println(string(encoding))
	count := 0
	for offset := 0; offset < 10000; offset += 100 {
		res, err := c.GetOwnershipDecentralizationHelper(ctx, offset)
		if err != nil {
			fmt.Println(err)
			break
		}
		for i := 0; i < 100; i++ {
			count++
			var r models.OwnershipDecentralization
			r.Address = res.Data[i].Address
			r.Balance = res.Data[i].Balance
			r.Blockchain = "Bitcoin"
			r.TimeStamp = Time
			data, _ := json.Marshal(r)
			fmt.Println(string(data))
			output := fmt.Sprintf("data/ownership_decentralization/bitcoin/%s%s.json", res.Data[i].Address, string(encoding))
			// fmt.Println(output)
			os.WriteFile(output, data, 0644)
		}
		time.Sleep(1 * time.Second)
	}
	fmt.Println(count)
}

func (c *Client) GetOwnershipDecentralizationHelper(ctx context.Context, offset int) (*Owner, error) {
    // limit := 100
    // page := 1
    // if options != nil {
    //     limit = options.Limit
    //     page = options.Page
    // }

    req, err := http.NewRequest("GET", fmt.Sprintf("%s%d",c.BaseURL, offset), nil)
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

