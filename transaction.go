package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"
)

type Response struct {
	Status       string `json:"status"`
	Transactions []struct {
		Type    string `json:"type"`

		Channel Reference

		Source Reference

		Paymentdate     time.Time `json:"paymentDate"`
		Sentdate        time.Time `json:"sentDate"`

		DeliveryAddress Address
		Billingaddress Address
		Senderaddress Address

		Quantity            int     `json:"quantity"`
		Productidentifier   string  `json:"productIdentifier"`
		Description         string  `json:"description"`
		Transactioncurrency string  `json:"transactionCurrency"`
		ItemPrice           float64 `json:"itemPrice"`
	} `json:"transactions"`
}

type Address struct {

	Fullname string `json:"fullName"`
	Street   string `json:"street"`
	Zip      string `json:"zip"`
	City     string `json:"city"`
	State    string `json:"state"`
	Country  string `json:"country"`
}

type Reference struct {

	Identifier        string `json:"identifier"`
	Transactionnumber string `json:"transactionNumber"`
	Itemnumber        string `json:"itemNumber"`
}

type TransactionsSummaryResult struct {

	Result 		string 										`json:"Result"`
	TransactionsTypePriceSums []TransactionsTypePriceSum	`json:"TransactionsTypePriceSums"`
}

type TransactionsTypePriceSum struct {
	Type 		string  	`json:"Type"`
	PriceSum 	float64  	`json:"PriceSum"`
}

type TransactionCreateRequest struct {

	Type		string 	`json:"type"`
	Quantity 	int 	`json:"quantity"`
}

type CreateTransactionResult struct {

	Result string `json:"Result"`
	Transaction TransactionCreateRequest  `json:"Transaction"`
}

type CreateTransactionBatchResult struct {

	CreateTransactionResults []CreateTransactionResult	`json:"CreateTransactionResults"`
}

func getTransactionsSummary() TransactionsSummaryResult{
	response, err := http.Get(config.DemoAPIWithJavaURL + "/api/transactions/")
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	transactionsTypePriceSum := make(map[string]float64)

	for i := 0; i < len(responseObject.Transactions); i++ {
		var transaction = responseObject.Transactions[i];

		transactionsTypePriceSum[ transaction.Type ] += transaction.ItemPrice
	}

	printLine()

	var result TransactionsSummaryResult

	result.Result = fmt.Sprintf("Total transactions count: %d", len(responseObject.Transactions))

	log.Println(result.Result);

	for key, value := range transactionsTypePriceSum {

		var typePriceSum TransactionsTypePriceSum
		typePriceSum.Type = key;
		typePriceSum.PriceSum = value;

		result.TransactionsTypePriceSums  = append(result.TransactionsTypePriceSums, typePriceSum)
		log.Printf("Type [%s] Item price sum = %.2f\n", key, value)
	}

	printLine()

	return result;
}

func createTransaction(index int) CreateTransactionResult {

	types := [2]string{"Sale", "Refund"}

	var request TransactionCreateRequest

	request.Type = types[rand.Intn(2)]
	request.Quantity = rand.Intn(9999)

	json_data, err := json.Marshal(request)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(config.DemoAPIWithJavaURL + "/api/transactions/add", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	json_request, err := json.Marshal(request)

	var result CreateTransactionResult

	result.Result = fmt.Sprintf("[%d] Created new transaction.", index);
	result.Transaction = request;

	log.Printf("%s %s\n", result.Result, json_request)

	return result
}

func createTransactionBatch(transactionCount int) CreateTransactionBatchResult {
	printLine()

	var result CreateTransactionBatchResult

	var wg sync.WaitGroup
	wg.Add(transactionCount)

	for i := 0; i < transactionCount; i++ {
		go func(i int) {
			defer wg.Done()

			result.CreateTransactionResults = append(result.CreateTransactionResults, createTransaction(i))
		}(i)
	}

	wg.Wait()

	log.Printf("\nCreated %d new transactions.\n", transactionCount)
	printLine()

	return result
}
