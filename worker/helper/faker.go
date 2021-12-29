package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/victorlau1/worker/models"
)

func CreateFakeNodeClientData(iterations int) {
	gofakeit.Seed(time.Now().Unix())
	b := new(bytes.Buffer)
	encoder := json.NewEncoder(b)

	timestamps := []string{
		time.Now().Add(time.Hour * -24 * 1).Format(time.RFC3339Nano),
		time.Now().Add(time.Hour * -24 * 2).Format(time.RFC3339Nano),
		time.Now().Add(time.Hour * -24 * 3).Format(time.RFC3339Nano),
		time.Now().Add(time.Hour * -24 * 4).Format(time.RFC3339Nano),
		time.Now().Add(time.Hour * -24 * 5).Format(time.RFC3339Nano),
		time.Now().Add(time.Hour * -24 * 6).Format(time.RFC3339Nano),
		time.Now().Add(time.Hour * -24 * 7).Format(time.RFC3339Nano),
	}

	for i := 1; i <= iterations; i++ {
		// fmt.Println(i % 6)

		nm := &models.ClientDecentralization{
			Timestamp: timestamps[i%6],
		}

		es := elasticIndex{
			Index: IndexAction{
				Index: indexName,
				Id:    fmt.Sprintf("%v", i),
			},
		}

		gofakeit.Struct(nm)
		encoder.Encode(es)
		encoder.Encode(nm)
	}
	dir := "/home/evorun/workspace/decentralized_dashboard/data/elastic_bulk_transformed"
	rdir, _ := os.ReadDir(dir)
	nf := fmt.Sprintf("%v/%v_%v.json", dir, "fake_data", len(rdir))
	nb, _ := ioutil.ReadAll(b)
	err := os.WriteFile(nf, nb, 0644)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}
