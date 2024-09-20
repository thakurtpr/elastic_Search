package connection

import (
	"fmt"
	v8 "github.com/elastic/go-elasticsearch/v8"
	// v7 "github.com/elastic/go-elasticsearch/v7"
)

var Es *v8.Client
var err error

func init() {
	cfg := v8.Config{
		Addresses: []string{
			"http://34.93.102.191:9200",
		},
		Username: "",
		Password: "",
	}

	Es, err = v8.NewClient(cfg)
	if err != nil {
		fmt.Println("Client initialisation error:",err)
	}

}
