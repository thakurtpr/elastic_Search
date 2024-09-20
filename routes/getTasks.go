package routes

import (
	"fmt"
	"net/http"
)

func GetTasks(response http.ResponseWriter,request *http.Request){
	fmt.Println("inside Get Tasks")

}