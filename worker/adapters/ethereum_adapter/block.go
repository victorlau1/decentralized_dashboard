package ethereum_adapter

import (
	// "context"
	"fmt"
	"context"
	// "net/url"
	"net/http"
	// "bytes"
	// "io"
	// "errors"
	"encoding/json"
	"os"
	"strconv"
	"github.com/victorlau1/worker/models"
	// "github.com/spf13/viper"
)

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


func (c *Client) GetBlockDecentralization(ctx context.Context, blockNumber int) (*models.BlockDecentralization, error) {

    req, err := http.NewRequest("GET", fmt.Sprintf("%s=%d&apiKey=%s", c.BaseURL, blockNumber, c.apiKey), nil)
    if err != nil {
        return nil, err
    }

    req = req.WithContext(ctx)

    res := BlockStats{}
    if err := c.sendRequest(req, &res); err != nil {
        return nil, err
    }
	data, _ := json.Marshal(res)
	os.WriteFile(fmt.Sprintf("raw_data/ethereum_block_%d.json", blockNumber), data, 0644)
    var r models.BlockDecentralization
    r.BlockNumber, _ = strconv.Atoi(res.Result.BlockNumber)
    r.TimeStamp, _ = strconv.Atoi(res.Result.TimeStamp)
    r.BlockMiner = res.Result.BlockMiner
    r.Blockchain = "Ethereum"
    return &r, nil
}
