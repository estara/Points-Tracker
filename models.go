package main

import "time"

type Transaction struct {
	payer     string    `json:"payer"`
	points    int       `json:"points"`
	timestamp time.Time `json:"timestamp"`
}

type Request struct {
	Points int `json:"points"`
}

var transactionList []Transaction

func getAllTransactions() []Transaction {
	return transactionList
}

func addTransaction(transaction) string {
	append(transactionList, transaction)
	return "added"
}
