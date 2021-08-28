package main

import "time"

type Transaction struct {
	Id        int       `json:"id"`
	Payer     string    `json:"payer"`
	Points    int       `json:"points"`
	Timestamp time.Time `json:"timestamp"`
	Used      int       `json:"used"`
}

type Request struct {
	Points int `json:"points"`
}

type Result struct {
	Payer  string `json:"payer"`
	Points int    `json:"points"`
}

var transactionList = []Transaction{Id: 1, Payer: "foo", Points: 5, Timestamp: "2020-11-02T14:00:00Z", Used: 0}

func getAllTransactions() []Transaction {
	return transactionList
}

func addTransaction(entry Transaction) {
	transactionList = append(transactionList, entry)
}
