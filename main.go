package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

var balance int

func main() {
    balance = 100

    r := gin.Default()

    r.POST("/increase", func(c *gin.Context) {
        amount := c.PostForm("amount")
        if num, err := strconv.Atoi(amount); err == nil {
            balance += num
            c.String(http.StatusOK, "Balance increased by %d. New balance: %d", num, balance)
        } else {
            c.String(http.StatusBadRequest, "Invalid amount")
        }
    })

    r.POST("/decrease", func(c *gin.Context) {
        amount := c.PostForm("amount")
        if num, err := strconv.Atoi(amount); err == nil {
            if balance >= num {
                balance -= num
                c.String(http.StatusOK, "Balance decreased by %d. New balance: %d", num, balance)
            } else {
                c.String(http.StatusBadRequest, "Not enough balance")
            }
        } else {
            c.String(http.StatusBadRequest, "Invalid amount")
        }
    })

    r.POST("/set", func(c *gin.Context) {
        amount := c.PostForm("amount")
        if num, err := strconv.Atoi(amount); err == nil {
            balance = num
            c.String(http.StatusOK, "Balance set to %d", balance)
        } else {
            c.String(http.StatusBadRequest, "Invalid amount")
        }
    })

    r.Run(":8080")
}

