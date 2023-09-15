package main

import (
	"fmt"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// uuid -> receipt
var receipts map[string]Receipt

/*
Originally was going to use a map to store the points for each receipt to avoid duplicate work.
Evidently this implementation would be beneficial for test cases where multiple receipts are being repeatedly processed.
*/
var Points = make(map[string]int)

var Items = []Item{
	{ShortDescription: "Milk", Price: "2.99"},
	{ShortDescription: "Pepsi - 12-oz", Price: "1.25"},
}

func main() {
	receipts = make(map[string]Receipt)

	router := gin.Default()
	router.GET("/Items", getItems)
	router.GET("/receipts", getReceipts)
	router.POST("/receipts/process", getReceipt)
	router.GET("/receipts/:id/points", receiptProcessor)
	router.Run("localhost:8087")
}

func receiptProcessor(c *gin.Context) {
	ID := c.Param("id")
	for k, v := range receipts {
		if k == ID {
			points := awardPoints(v)
			// Avoid duplicate processing
			Points[ID] = points
			c.IndentedJSON(http.StatusOK, gin.H{"points": points})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Receipt not found"})
}

// Use provided rules to compute points for a given receipt
func awardPoints(r Receipt) int {
	points := 0
	points += alphanumericCharaterPoints(r.Retailer)
	points += roundDollarAmmount(r.Total)
	points += multipleofQuarter(r.Total)
	points += everyTwoItems(r.Items)
	points += trimmedLength(r.Items)
	points += oddDay(r.PurchaseDate)
	points += happyhour(r.PurchaseTime)
	return points
}

func alphanumericCharaterPoints(s string) int {
	return len(regexp.MustCompile("[^a-zA-Z0-9]+").ReplaceAllString(s, ""))
}

func roundDollarAmmount(total string) int {
	points, err := strconv.ParseFloat(total, 64)
	// Invalid format
	if err != nil {
		return 0
	}

	// Check if there are no cents
	if points == float64(int(points)) {
		return 50
	}
	return 0
}

func multipleofQuarter(total string) int {
	points, err := strconv.ParseFloat(total, 64)
	// Invalid format
	if err != nil {
		return 0
	}

	// Check if total is multiple of 0.25
	if math.Mod(points, 0.25) == 0 {
		return 25
	}
	return 0
}

func everyTwoItems(items []Item) int {
	return (len(items) / 2) * 5
}

/*
If the trimmed length of the item description is a multiple of 3,
multiply the price by 0.2 and round up to the nearest integer.
The result is the number of points earned.
*/
func trimmedLength(items []Item) int {
	points := 0
	for i := 0; i < len(items); i++ {
		item := items[i]
		trimmedDescription := strings.Trim(item.ShortDescription, " ")

		if len(trimmedDescription)%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err == nil {
				points += int(math.Ceil(price * 0.2))
			}
		}
	}
	return points
}

func oddDay(date string) int {
	purchaseDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		fmt.Printf("Error parsing purchase date for item: %v\n", err)
	}
	if purchaseDate.Day()%2 == 1 {
		return 6
	}
	return 0
}

// 10 points if the time of purchase is after 2:00pm and before 4:00pm
func happyhour(purchaseTime string) int {
	hourString := strings.Split(purchaseTime, ":")[0]
	hour, err := strconv.Atoi(hourString)
	if err != nil {
		return 0
	}
	if hour >= 14 && hour < 16 {
		//fmt.Println("Happy hour!")
		return 10
	}
	return 0
}

func getReceipts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, receipts)
}

func createUUID() string {
	return uuid.New().String()
}

// getItems responds with the list of all items as JSON.
func getItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Items)
}

// Submits a receipt for processing -> Returns the ID assigned to the receipt
func getReceipt(c *gin.Context) {
	var receipt Receipt
	err := c.BindJSON(&receipt)
	// Call BindJSON to bind the received JSON to receipt.
	if err != nil {
		return
	}
	id := createUUID()
	receipts[id] = receipt
	c.IndentedJSON(http.StatusOK, gin.H{"id": id})
}

/* Model */

// Items struct to be used in the Receipt struct
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Total        string `json:"total"`
	Items        []Item `json:"items"`
}
