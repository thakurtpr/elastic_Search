package routes

import (
	"elastic_Search/connection"
	"fmt"
	"net/http"
	"strings"
)

func GetTasks(response http.ResponseWriter, request *http.Request) {
    fmt.Println("inside Get Tasks")


    if connection.Es == nil {
        http.Error(response, "Elasticsearch client not initialized", http.StatusInternalServerError)
        return
    }

    body := `{
        "query": {
            "match": { "assignee": "ajit" }
        }
    }`

    res, err := connection.Es.Search(
        connection.Es.Search.WithIndex("*"),
        connection.Es.Search.WithBody(strings.NewReader(body)),
        connection.Es.Search.WithPretty(),
    )
    if err != nil {
        http.Error(response, fmt.Sprintf("Error searching Elasticsearch: %v", err), http.StatusInternalServerError)
        return
    }

	
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(res.String()))
    // Optionally, you can handle the response further or write it to the response writer
    // fmt.Println(res)
}
