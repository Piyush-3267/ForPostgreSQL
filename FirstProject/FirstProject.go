package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "DemoTitle", Desc: "DemoDesc", Content: "DemoContent"},
	}
	fmt.Fprintf(w, "Welcome to the HomePage! ")
	fmt.Println("\n Endpoint Hit: homePage")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func ForGorilaMuxGET(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "we are because of get Method")

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/article", ForGorilaMuxGET).Methods("GET")
	myRouter.HandleFunc("/article", allArticles).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()
}
