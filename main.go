package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/balances", getBalances)
	router.POST("/spend", spendPoints)
	router.POST("/transactions", addTransaction)
	// store := cookie.NewStore([]byte("secret"))
	// router.Use(sessions.Sessions("mysession", store))

	// router.GET("/incr", func(context *gin.Context) {
	// 	session := sessions.Default(context)
	// 	var count int
	// 	v := session.Get("count")
	// 	if v == nil {
	// 		count = 0
	// 	} else {
	// 		count = v.(int)
	// 		count++
	// 	}
	// 	session.Set("count", count)
	// 	session.Save()
	// 	context.JSON(200, gin.H{"count": count})
	// })

	router.Run(":8080")
}

// var globalSessions *session.Manager

// // initialize in init() function
// func init() {
// 	globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid","gclifetime":3600}`)
// 	go globalSessions.GC()
// }
