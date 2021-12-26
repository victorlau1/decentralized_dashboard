package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type BitNode struct {
	Timestamp    int                      `json:"timestamp"`
	TotalNodes   int                      `json:"total_nodes"`
	LatestHeight int                      `json:"latest_height"`
	Nodes        map[string][]interface{} `json:"nodes"`
}

type NodeValues struct {
	IPAddress        string
	ProtocolVersion  float64
	UserAgent        string
	ConnectedSince   float64
	Services         float64
	Height           float64
	Hostname         string
	City             string
	CountryCode      string
	Latitude         float64
	Longitude        float64
	Timezone         string
	ASN              string
	OrganizationName string
	Timestamp        time.Time
}

func JSONToNewLineJSON() {
	filename := "/home/evorun/workspace/decentralized_dashboard/data/2021_12_22_8_56_nodes_data.json"
	f, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error %v", err)
	}

	// fmt.Printf("%v", f)

	bitNodeResponse := BitNode{}
	err = json.Unmarshal(f, &bitNodeResponse)

	if err != nil {
		fmt.Printf("Error %v", err)
	}

	b := new(bytes.Buffer)
	enc := json.NewEncoder(b)
	for key, element := range bitNodeResponse.Nodes {
		nn := NodeValues{}
		nn.IPAddress = key
		nn.ProtocolVersion = element[0].(float64)
		if ua, ok := element[1].(string); ok {
			nn.UserAgent = ua
		}
		nn.ConnectedSince = element[2].(float64)
		nn.Services = element[3].(float64)
		nn.Height = element[4].(float64)
		if hName, ok := element[5].(string); ok {
			nn.Hostname = hName
		}

		if city, ok := element[6].(string); ok {
			nn.City = city
		}
		if country, ok := element[7].(string); ok {
			nn.CountryCode = country
		}
		nn.Latitude = element[8].(float64)
		nn.Longitude = element[9].(float64)
		if timeZone, ok := element[10].(string); ok {
			nn.Timezone = timeZone
		}
		nn.ASN = element[11].(string)
		if orgName, ok := element[12].(string); ok {
			nn.OrganizationName = orgName
		}

		enc.Encode(nn)
	}
	current_timestamp := time.Now()
	nf := fmt.Sprintf("/home/evorun/workspace/decentralized_dashboard/data/transformed/bit_nodes%v.json", current_timestamp)
	nb, _ := ioutil.ReadAll(b)
	os.WriteFile(nf, nb, 0644)

}

func SolanaToNewLineJSON() {

}
