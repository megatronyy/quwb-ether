package web

import (
	"net/http"
	"fmt"
)

func Run(){
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8085", nil)
}

func hello(response http.ResponseWriter, request *http.Request){
	fmt.Fprintf(response, "Hello, Welcome to go web programming...")
}
