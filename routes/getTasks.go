package routes

import (
	"elastic_Search/connection"
	"fmt"
	"net/http"
	"strings"
)

func GetTasks(response http.ResponseWriter,request *http.Request){
	fmt.Println("inside Get Tasks")
	body := `{
		"query": {
			"match": { "assignee": "demo" }
		},
	}`
	res, err := connection.Es.Search(
		connection.Es.Search.WithIndex("*"),
		connection.Es.Search.WithBody(strings.NewReader(body)),
		connection.Es.Search.WithPretty(),
	)
	fmt.Println(res,err)
}