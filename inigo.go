package main

import (
"fmt"
"net/http")

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res,"Hello my name is ...1111")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("0.0.0.0:80", nil)
}