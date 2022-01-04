package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/pariz/gountries"
	sw "github.com/victorlau1/solanaclient"
	sol "github.com/victorlau1/worker/adapters/solana_adapter"
	"github.com/victorlau1/worker/models"
)

var fakeTimestamp = time.Now().Add(time.Hour * 24 * -1)
var bitnodeFilename = "/home/evorun/workspace/decentralized_dashboard/data/2021_12_23_21_52_bitcoin_nodes_data.json"
var ethereumFilename = "/home/evorun/workspace/decentralized_dashboard/data/2021_12_25_ethernodes.txt"
var solanaValidatorFilename = "/home/evorun/workspace/decentralized_dashboard/data/2021_12_23_21_54_solana_validators.json"
var solanaNonValidatorFilename = "/home/evorun/workspace/decentralized_dashboard/data/solana_node_nonvalidators_data_2021-12-25T16:53:57-08:00.json"

type BitNode struct {
	Timestamp    int                      `json:"timestamp"`
	TotalNodes   int                      `json:"total_nodes"`
	LatestHeight int                      `json:"latest_height"`
	Nodes        map[string][]interface{} `json:"nodes"`
}

type NodeValues struct {
	IPAddress        string  //0 -- this is the key
	ProtocolVersion  float64 //1
	UserAgent        string  //2
	ConnectedSince   float64 //3
	Services         float64 //4
	Height           float64 //5
	Hostname         string  //6
	City             string  //7
	CountryCode      string  //8
	Latitude         float64 //9
	Longitude        float64 //10
	Timezone         string  //11
	ASN              string  //12
	OrganizationName string
	Timestamp        time.Time
}

func BitNodeToNewLineJSON() {
	f, err := os.ReadFile(bitnodeFilename)
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

		if country, ok := element[7].(string); ok {
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

		asn, _ := strconv.ParseFloat(strings.Replace(element[11].(string), "AS", "", 1), 64)
		nm.ASN = asn

		if orgName, ok := element[12].(string); ok {
			nm.Organization = orgName
		}

		nm.Timestamp = fakeTimestamp.Format(time.RFC3339Nano)

		enc2.Encode(nm)
	}
	current_timestamp := fakeTimestamp.Format("20060102150405")
	nf := fmt.Sprintf("/home/evorun/workspace/decentralized_dashboard/data/transformed/bit_nodes_%v.json", current_timestamp)
	nb, _ := ioutil.ReadAll(b2)
	os.WriteFile(nf, nb, 0644)
}

func SolanaToNewLineJSON() {
	f, err := os.ReadFile(solanaNonValidatorFilename)
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

		nm.Client = "Solana"

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
		nm.Timestamp = fakeTimestamp.Format(time.RFC3339Nano) //Used to fake timestamp data
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

	current_timestamp := fakeTimestamp.Format("20060102150405") //Used to fake timestamp data
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
	f, err := os.ReadFile(solanaValidatorFilename)
	if err != nil {
		fmt.Printf("Error %v", err)
	}

	valResp := &[]SolValidators{}
	json.Unmarshal(f, valResp)

	b := new(bytes.Buffer)

	enc := json.NewEncoder(b)

	for _, val := range *valResp {
		nm := models.ClientDecentralization{}

		client := sol.NewClient(nil, nil)
		vac := client.GetValidatorNode(val.VoteAccount)

		fmt.Println(val.VoteAccount)
		nm.ASN = float64(val.AutonomousSystemNumber)
		nm.Blockchain = "Solana"

		if vac.Validator.GetLocation().City != nil {
			nm.City = (*vac.Validator.Location.City)
		}
		if vac.Validator.GetLocation().Country != nil {
			nm.Country = (*vac.Validator.Location.Country)
		}

		nm.Client = "Solana"
		nm.ClientVersion = val.SoftwareVersion

		// nm.IPAddress =
		if vac.Validator.GetLocation().Ll != nil {
			nm.Latitude = float64((*vac.Validator.Location.Ll)[0])
			nm.Longitude = float64((*vac.Validator.Location.Ll)[1])
		}

		if vac.Validator.GetAsn().Organization != nil {
			nm.Organization = (*vac.Validator.Asn.Organization)
		}

		if vac.Validator.GetLocation().Region != nil {
			nm.Region = (*vac.Validator.Location.Region)
		}

		nm.Validator = true
		nm.Timestamp = fakeTimestamp.Format(time.RFC3339Nano)

		if vac.Validator.GetLocation().Timezone != nil {
			nm.Timezone = (*vac.Validator.Location.Timezone)
		}

		enc.Encode(nm)
		time.Sleep(1 * time.Microsecond) //100 requests per second so this throttles
	}

	current_timestamp := fakeTimestamp.Format("20060102150405")
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
	f, err := os.ReadFile(ethereumFilename)
	if err != nil {
		fmt.Printf("Error %v", err)
	}

	b := new(bytes.Buffer)
	enc := json.NewEncoder(b)

	query := gountries.New()

	data := string(f)
	rows := strings.Split(data, "\n")
	for _, row := range rows {
		cols := strings.Split(row, "\t")
		// fmt.Println(cols)
		nm := models.ClientDecentralization{}

		country, _ := query.FindCountryByName(cols[3])

		nm.Country = country.Alpha2
		nm.Client = cols[4]
		//nm.Region
		//nm.Timezone
		//nm.City
		nm.Blockchain = "Ethereum"
		nm.Timestamp = fakeTimestamp.Format(time.RFC3339Nano)
		nm.ClientVersion = cols[5]
		nm.Organization = cols[2] //Host is equivalent to organization?
		nm.IPAddress = cols[1]
		//nm.Longitude
		//nm.Latitude
		//nm.ASN
		enc.Encode(nm)
	}

	current_timestamp := fakeTimestamp.Format("20060102150405")
	nf := fmt.Sprintf("/home/evorun/workspace/decentralized_dashboard/data/transformed/ethereum_client_%v.json", current_timestamp)
	nb, _ := ioutil.ReadAll(b)
	os.WriteFile(nf, nb, 0644)
}
