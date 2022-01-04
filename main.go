package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Desc", Content: "Test Content"},
	}

	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Test Post Articles Endpoint")
	fmt.Println("Method:", r.Method)
	fmt.Println("Path:", r.URL.Path)
	fmt.Println("Authorization:", r.Header.Get("Authorization"))
	fmt.Fprintf(w, "Post endpoint hit")

	// Get the request body and decode into struct
	reqBody := &Article{}
	err := json.NewDecoder(r.Body).Decode(reqBody)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(reqBody.Title)
	fmt.Println(reqBody.Desc)
	fmt.Println(reqBody.Content)
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequest()
}
