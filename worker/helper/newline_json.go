package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	sw "github.com/victorlau1/solanaclient"
	"github.com/victorlau1/worker/models"
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

func BitNodeToNewLineJSON() {
	filename := "/home/evorun/workspace/decentralized_dashboard/data/2021_12_22_8_56_bitcoin_nodes_data.json"
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
	b2 := new(bytes.Buffer)
	enc := json.NewEncoder(b)
	enc2 := json.NewEncoder(b2)
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

		//Same schema

		nm := models.ClientDecentralization{}

		if country, ok := element[6].(string); ok {
			nm.Country = country
		}

		if ua, ok := element[1].(string); ok {
			nm.Client = ua
		}

		if timeZone, ok := element[10].(string); ok {
			nm.Timezone = timeZone
		}

		nm.Blockchain = "Bitcoin"

		nm.IPAddress = key

		nm.Latitude = element[8].(float64)
		nm.Longitude = element[9].(float64)

		if orgName, ok := element[12].(string); ok {
			nm.Organization = orgName
		}

		enc2.Encode(nm)
	}
	current_timestamp := time.Now()
	nf := fmt.Sprintf("/home/evorun/workspace/decentralized_dashboard/data/transformed/bit_nodes%v.json", current_timestamp)
	nb, _ := ioutil.ReadAll(b2)
	os.WriteFile(nf, nb, 0644)
}

func SolanaToNewLineJSON() {
	filename := "/home/evorun/workspace/decentralized_dashboard/data/solana_node_nonvalidators_data_2021-12-22T21:31:40-08:00.json"
	f, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error %v", err)
	}

	valResp := &[]sw.InlineResponse20015{}

	b := new(bytes.Buffer)
	enc := json.NewEncoder(b)

	json.Unmarshal(f, valResp)

	for _, v := range *valResp {
		nm := &models.ClientDecentralization{}

		loc := v.GetLocation()

		if loc.Country != nil {
			nm.Country = (*loc.Country)
		}

		nm.Client = "RustClient"

		if loc.Region != nil {
			nm.Region = (*v.GetLocation().Region)
		}

		if loc.Timezone != nil {
			nm.Timezone = (*v.GetLocation().Timezone)
		}

		if loc.City != nil {
			nm.City = (*v.GetLocation().City)
		}

		nm.Blockchain = "solana"
		nm.Timestamp = time.Now()
		nm.ClientVersion = v.GetVersion()

		if v.GetAsn().Organization != nil {
			nm.Organization = (*v.GetAsn().Organization)
		}

		//	nm.IPAddress = Doesn't Exist for NonValidators it seems
		if v.GetLocation().Ll != nil {
			nm.Latitude = float64((*v.GetLocation().Ll)[0])
			nm.Longitude = float64((*v.GetLocation().Ll)[1])
		}

		if v.GetAsn().Code != nil {
			nm.ASN = float64((*v.GetAsn().Code))
		}

		enc.Encode(nm)
	}

	current_timestamp := time.Now()
	nf := fmt.Sprintf("/home/evorun/workspace/decentralized_dashboard/data/transformed/non_val_nodes_%v.json", current_timestamp)
	nb, _ := ioutil.ReadAll(b)
	os.WriteFile(nf, nb, 0644)
}

type SolValidators struct {
	Network                      string    `json:"network"`
	Account                      string    `json:"account"`
	Name                         string    `json:"name"`
	KeybaseId                    string    `json:"keybase_id"`
	WwwUrl                       string    `json:"www_url"`
	Details                      string    `json:"details"`
	AvatarUrl                    string    `json:"avatar_url"`
	CreatedAt                    time.Time `json:"created_at"`
	UpdatedAt                    time.Time `json:"updated_at"`
	TotalScore                   int       `json:"total_score"`
	RootDistanceScore            int       `json:"root_distance_score"`
	VoteDistanceScore            int       `json:"vote_distance_score"`
	SkippedSlotScore             int       `json:"skipped_slot_score"`
	SoftwareVersion              string    `json:"software_version"`
	SoftwareVersionScore         int       `json:"software_version_score"`
	StakeConcentrationScore      int       `json:"stake_concentration_score"`
	DataCenterConcentrationScore int       `json:"data_center_concentration_score"`
	PublishedInformationScore    int       `json:"published_information_score"`
	SecurityReportScore          int       `json:"security_report_score"`
	ActiveStake                  int       `json:"active_stake"`
	Commission                   int       `json:"commission"`
	Delinquent                   bool      `json:"delinquent"`
	DataCenterKey                string    `json:"data_center_key"`
	DataCenterHost               string    `json:"data_center_host"`
	AutonomousSystemNumber       int       `json:"autonomous_system_number"`
	VoteAccount                  string    `json:"vote_account"` //Pubkey
	EpochCredits                 int       `json:"epoch_credits"`
	SkippedSlots                 int       `json:"skipped_slots"`
	SkippedSlotPercent           string    `json:"skipped_slot_percent"`
	PingTime                     string    `json:"ping_time"`
	Url                          string    `json:"url"`
}

func SolanaValidatorsToNewLineJSON() {
	b := new(bytes.Buffer)

	current_timestamp := time.Now()
	nf := fmt.Sprintf("/home/evorun/workspace/decentralized_dashboard/data/transformed/val_nodes_%v.json", current_timestamp)
	nb, _ := ioutil.ReadAll(b)
	os.WriteFile(nf, nb, 0644)
}

type EthereumNodes struct {
	Id       string
	Host     string
	ISP      string
	Country  string
	Client   string
	Version  string
	OS       string
	LastSeen string
	InSync   bool
}

func EthereumToNewLineJSON() {
	filename := "/home/evorun/workspace/decentralized_dashboard/data/2021_12_25_ethernodes.txt"
	f, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error %v", err)
	}

	b := new(bytes.Buffer)
	enc := json.NewEncoder(b)

	data := string(f)
	rows := strings.Split(data, "\n")
	for _, row := range rows {
		cols := strings.Split(row, "\t")
		// fmt.Println(cols)
		nm := models.ClientDecentralization{}
		nm.Country = cols[3]
		nm.Client = cols[4]
		//nm.Region
		//nm.Timezone
		//nm.City
		nm.Blockchain = "Ethereum"
		nm.Timestamp = time.Now()
		nm.ClientVersion = cols[5]
		//nm.Organization
		nm.IPAddress = cols[1]
		//nm.Longitude
		//nm.Latitude
		//nm.ASN
		enc.Encode(nm)
	}

	current_timestamp := time.Now()
	nf := fmt.Sprintf("/home/evorun/workspace/decentralized_dashboard/data/transformed/ethereum_%v.json", current_timestamp)
	nb, _ := ioutil.ReadAll(b)
	os.WriteFile(nf, nb, 0644)
}
