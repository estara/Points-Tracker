package main

import (
	"sort"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func pointCount(transactions []Transaction) map[string]int {
	points := make(map[string]int)
	for _, transaction := range transactions {
		_, ok := points[trasaction.payer]
		if ok == true {
			points[transaction.payer] += transaction.points
		} else {
			points[transaction.payer] = transaction.points
		}
	}
	return points
}

func spendPoints(c *gin.Context) {
	var newRequest Request
	if err := c.BindJSON(&newRequest); err != nil {
        return
    }
	transactions := getAllTransactions()
	for _, transaction := range transactions {

	}
	if balance < newRequest.points {
		c.JSON(400, "Insufficient points")
	}
	sort.Slice(transactions[:], func(i, j int) bool {
		return transactions[i].timestamp < transactions[j].timestamp
	  })
	
	// Call the JSON method of the Context to return a message.
	c.JSON(http.StatusOK, gin.H{ 
		"code" : http.StatusOK, 
		"message": string({"payer": payer, "points": points}),// cast it to string before showing
})

}

func getBalances(c *gin.Context) {
	transactions = getAllTransactions()
	const balances := pointCount(transactions)
	c.JSON(http.StatusOK, string(balances))
}

func addTransaction(c *gin.Context) {
    var newTransaction Transaction
    if err := c.BindJSON(&newTransaction); err != nil {
        return
    }
	newTransaction.timestamp = time.Parse(time.RFC3339, newTransaction.timestamp)
    addTransaction(newTransaction)
    c.JSON(http.StatusCreated, newTransaction)
}