package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
	"routes"
)

func main(){
	err := godotenv.Load("./env/.config")
	if err != nil {
		log.Fatal("Error loading .env file :", err)
	}
	
	fmt.Println(os.Getenv("ELASTIC_URL"))

	routes:=mux.NewRouter();
	routes.Get("/",getTasks)
}