package main

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

// count how many points each payer has
func pointCount(transactions []Transaction) map[string]int {
	points := make(map[string]int)
	for _, transaction := range transactions {
		_, ok := points[transaction.Payer]
		if ok == true {
			points[transaction.Payer] += transaction.Points
		} else {
			points[transaction.Payer] = transaction.Points
		}
	}
	return points
}

// POST request to spend points as { "points": 5000 }
// returns list of payers whose points were spent:
//[
//	{ "payer": "DANNON", "points": -100 },
// 	{ "payer": "UNILEVER", "points": -200 },
// 	{ "payer": "MILLER COORS", "points": -4,700 }
// ]
func spendPoints(c *gin.Context) {
	var newRequest Request
	if err := c.BindJSON(&newRequest); err != nil {
		return
	}

	var balance int = 0
	transactions := getAllTransactions()
	for _, transaction := range transactions {
		balance += transaction.Points
	}
	if balance < newRequest.Points {
		c.JSON(400, "Insufficient points")
	}

	sort.SliceStable(transactions[:], func(i, j int) bool {
		return transactions[i].Timestamp.Before(transactions[j].Timestamp)
	})
	var resultTransactions []Transaction
	var resultBalance int = 0
	for _, transaction := range transactions {
		if resultBalance < newRequest.Points {
			resultTransactions = append(resultTransactions, transaction)
			resultBalance += transaction.Points
		} else {
			break
		}
	}
	if resultBalance > newRequest.Points {
		resultTransactions[len(resultTransactions)-1].Points -= (resultBalance - newRequest.Points)
	}
	payerPoints := pointCount(resultTransactions)
	var pointsSpent []Result
	for payer, points := range payerPoints {
		var newResult = Result{Payer: payer, Points: -points}
		pointsSpent = append(pointsSpent, newResult)
	}
	c.JSON(http.StatusOK, pointsSpent)
}

// GET request, returns the current balance for all payers
//{
// "DANNON": 1000,
// "UNILEVER": 0,
// "MILLER COORS": 5300
// }
func getBalances(c *gin.Context) {
	transactions := getAllTransactions()
	balances := pointCount(transactions)
	c.JSON(http.StatusOK, balances)
}

// POST request to add a new transaction:
// { "payer": "DANNON", "points": 1000, "timestamp": "2020-11-02T14:00:00Z" }
// returns "added"
func newTransaction(c *gin.Context) {
	var newTransaction Transaction
	if err := c.BindJSON(&newTransaction); err != nil {
		return
	}
	addTransaction(newTransaction)
	c.JSON(http.StatusCreated, "added")
}
