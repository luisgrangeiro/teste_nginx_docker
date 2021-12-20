package main

import(
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", GetFunc)
	http.ListenAndServe(":9000", nil)
}

func GetFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello From Golang Api")
}