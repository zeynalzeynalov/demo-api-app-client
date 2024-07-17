package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func checkApiStatusPage(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w, "DemoApiClientApplication API is active.")
	log.Println("HTTP request to endpoint: checkApiStatusPage")
}

func createTransactionsPage(w http.ResponseWriter, r *http.Request){

	var result = createTransactionBatch(10)

	json, err := json.Marshal(result)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "%s", json)
	log.Println("HTTP request to endpoint: createTransactionsPage")
}

func getTransactionsSummaryPage(w http.ResponseWriter, r *http.Request){

	var result = getTransactionsSummary()

	json, err := json.Marshal(result)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "%s", json)
	log.Println("HTTP request to endpoint: getTransactionsSummaryPage")
}

func handleRequests() {

	port := fmt.Sprintf("%d", config.WebServerPort)

	const indexPage = "public/index.html"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if buf, err := ioutil.ReadAll(r.Body); err == nil {
				log.Printf("Received message: %s\n", string(buf))
			}
		} else {
			log.Printf("Serving %s to %s...\n", indexPage, r.RemoteAddr)
			http.ServeFile(w, r, indexPage)
		}
	})

	http.HandleFunc("/api/", checkApiStatusPage)
	http.HandleFunc("/api/transactions/generate", createTransactionsPage)
	http.HandleFunc("/api/transactions/summary", getTransactionsSummaryPage)

	log.Printf("Listening on port %s\n\n", port)

	log.Fatal(http.ListenAndServe(":" + port, nil))
}
