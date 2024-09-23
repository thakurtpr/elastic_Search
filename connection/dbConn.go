package connection

import (
	"fmt"
	"os"
	"log"

	v8 "github.com/elastic/go-elasticsearch/v8"
	// v7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/joho/godotenv"
)

var Es *v8.Client
var err error

func init() {
	err := godotenv.Load("./env/.config")
	if err != nil {
		log.Fatal("Error loading .env file :", err)
	}
	cfg := v8.Config{
		Addresses: []string{
			os.Getenv("ELASTIC_URL"),
		},
		Username: os.Getenv("USER_NAME"),
		Password: os.Getenv("PASSWORD"),
	}

	Es, err = v8.NewClient(cfg)
	if err != nil {
		fmt.Println("Client initialisation error:",err)
	}

}
