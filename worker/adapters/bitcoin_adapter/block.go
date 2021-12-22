package bitcoin_adapter

import (
	// "context"
	"fmt"
	"context"
	// "net/url"
	"net/http"
	// "bytes"
	// "io"
	"github.com/victorlau1/worker/models"
	"os"
	"encoding/json"
	// "github.com/spf13/viper"
)

type BlockStats struct {
	Blocks []struct {
		Hash       string   `json:"hash"`
		Ver        int      `json:"ver"`
		PrevBlock  string   `json:"prev_block"`
		MrklRoot   string   `json:"mrkl_root"`
		Time       int      `json:"time"`
		Bits       int      `json:"bits"`
		NextBlock  []string `json:"next_block"`
		Fee        int      `json:"fee"`
		Nonce      int      `json:"nonce"`
		NTx        int      `json:"n_tx"`
		Size       int      `json:"size"`
		BlockIndex int      `json:"block_index"`
		MainChain  bool     `json:"main_chain"`
		Height     int      `json:"height"`
		Weight     int      `json:"weight"`
		Tx         []struct {
			Hash        string `json:"hash"`
			Ver         int    `json:"ver"`
			VinSz       int    `json:"vin_sz"`
			VoutSz      int    `json:"vout_sz"`
			Size        int    `json:"size"`
			Weight      int    `json:"weight"`
			Fee         int    `json:"fee"`
			RelayedBy   string `json:"relayed_by"`
			LockTime    int    `json:"lock_time"`
			TxIndex     int64  `json:"tx_index"`
			DoubleSpend bool   `json:"double_spend"`
			Time        int    `json:"time"`
			BlockIndex  int    `json:"block_index"`
			BlockHeight int    `json:"block_height"`
			Inputs      []struct {
				Sequence int64  `json:"sequence"`
				Witness  string `json:"witness"`
				Script   string `json:"script"`
				Index    int    `json:"index"`
				PrevOut  struct {
					Spent             bool   `json:"spent"`
					Script            string `json:"script"`
					SpendingOutpoints []struct {
						TxIndex int64 `json:"tx_index"`
						N       int   `json:"n"`
					} `json:"spending_outpoints"`
					TxIndex int   `json:"tx_index"`
					Value   int   `json:"value"`
					N       int64 `json:"n"`
					Type    int   `json:"type"`
				} `json:"prev_out"`
			} `json:"inputs"`
			Out []struct {
				Type              int   `json:"type"`
				Spent             bool  `json:"spent"`
				Value             int64 `json:"value"`
				SpendingOutpoints []struct {
					TxIndex int64 `json:"tx_index"`
					N       int   `json:"n"`
				} `json:"spending_outpoints"`
				N       int    `json:"n"`
				TxIndex int64  `json:"tx_index"`
				Script  string `json:"script"`
				Addr    string `json:"addr"`
			} `json:"out"`
		} `json:"tx"`
	} `json:"blocks"`
}

func (c *Client) GetBlockDecentralization(ctx context.Context, blockNumber int) (*models.BlockDecentralization, error) {

    req, err := http.NewRequest("GET", fmt.Sprintf("%s/%d?format=json", c.BaseURL, blockNumber), nil)
    if err != nil {
        return nil, err
    }

    req = req.WithContext(ctx)

    res := BlockStats{}
    if err := c.sendRequest(req, &res); err != nil {
        return nil, err
    }
	data, _ := json.Marshal(res)
	os.WriteFile(fmt.Sprintf("raw_data/bitcoin_block_%d.json", blockNumber), data, 0644)
    var r models.BlockDecentralization
    r.BlockNumber = res.Blocks[0].Height
    r.TimeStamp = res.Blocks[0].Time
    r.BlockMiner = res.Blocks[0].Tx[0].Out[0].Addr
    r.Blockchain = "Bitcoin"
    return &r, nil
}