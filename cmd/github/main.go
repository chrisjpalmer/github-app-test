package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chrisjpalmer/github-app-test/internal/handler"
	"github.com/gorilla/mux"
)

const appID = int64(943300)

func main() {

	addr := ":8080"
	fmt.Println("listening on ", addr)

	h := handler.New(appID)

	// listen for webhook events
	r := mux.NewRouter()
	r.HandleFunc("/callback", h.Callback)
	err := http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatal(err)
	}
}
