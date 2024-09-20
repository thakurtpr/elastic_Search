package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"elastic_Search/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main(){
	err := godotenv.Load("./env/.config")
	if err != nil {
		log.Fatal("Error loading .env file :", err)
	}
	
	fmt.Println(os.Getenv("ELASTIC_URL"))

	r:=mux.NewRouter();
	r.HandleFunc("/",routes.GetTasks)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PATCH"},
		// AllowedMethods: []string{"*"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
		// AllowedHeaders: []string{"*"},
	})

	handler := c.Handler(r)
	port := ":4005"
	s := &http.Server{
		Addr:    port,
		Handler: handler,
	}

	log.Printf("Server is Running in Port %v", port)
	log.Fatal(s.ListenAndServe())
}