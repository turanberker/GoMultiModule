package main

import (
	"log"
	"net/http"

	"com.mbt.authz.common/service"
)

type apiHandler struct {
	id int
}

func (apiHandler) ServeHTTP(http.ResponseWriter, *http.Request) {}

func main() {

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		service.CommonServiceTest()
	}

	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
