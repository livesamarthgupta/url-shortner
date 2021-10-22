package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/urlshortner/database"
	"github.com/urlshortner/url"
)


const API_BASE_PATH="/api"
const PORT=":80"

// TODO: Run SQL script on first call

func main() {
	database.SetupDatabase()
	err := url.SetupTable()  
	if err != nil {
		log.Fatal(err)
	}
	url.SetupRoutes(API_BASE_PATH)
	http.ListenAndServe(PORT, nil)
	log.Println("Listening on port: " + PORT)
}