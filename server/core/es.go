package core

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"os"
	"server/global"
)

func ESInit() *elasticsearch.Client {
	//client, err := elasticsearch.NewClient(elasticsearch.Config{
	//	CloudID: "acc",
	//	APIKey:  "LUtwZUo0NEJJdEZjZExtYmxjQ2Y6cXd2XzI3ZnpSd0dmdFVDSzVYQVdZZw==",
	//})
	cert, err := os.ReadFile("D:\\es\\http_ca.crt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
		},
		Username: "elastic",
		Password: global.Config.Elasticsearch.Password,
		CACert:   cert,
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		fmt.Println("Error creating the client:", err)
	}

	return es
}
