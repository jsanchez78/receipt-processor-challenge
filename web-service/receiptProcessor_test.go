package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

// Parse JSON data
var receiptsTest []Receipt

func TestSetup(t *testing.T) {
	jsonFile, err := os.ReadFile("input.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Unmarshal data into Receipt struct
	err = json.Unmarshal(jsonFile, &receiptsTest)
	if err != nil {
		fmt.Println("Error parsing JSON data:", err)
		return
	}
}
func TestComputeTotalPoints(t *testing.T) {
	expectedResults := []int{28, 109}
	// Loop through each receipt object  and compute the total points
	for i, receipt := range receiptsTest {
		expected := expectedResults[i]
		result := awardPoints(receipt)
		if result != expected {
			t.Errorf("Expected %d points, got %d", expected, result)
		}
		fmt.Printf("Total points: %d\n", result)
	}
}
func TestAlphanumericCharaterPoints(t *testing.T) {
	expectedResults := []int{6, 14}
	// Loop through each receipt object  and compute the total points
	for i, receipt := range receiptsTest {
		expected := expectedResults[i]
		result := alphanumericCharaterPoints(receipt.Retailer)
		if result != expected {
			t.Errorf("Expected %d points, got %d", expected, result)
		}
		//fmt.Printf("\n%d points - retailer name %s\n"+
		//"has %d alphanumeric characters note: '&' is not alphanumeric\n", result, receipt.Retailer, result)
	}
}
func TestRoundDollarAmmount(t *testing.T) {
	expectedResults := []int{0, 50}
	// Loop through each receipt object  and compute the total points
	for i, receipt := range receiptsTest {
		expected := expectedResults[i]
		result := roundDollarAmmount(receipt.Total)
		if result != expected {
			t.Errorf("Expected %d points, got %d", expected, result)
		}
	}
}
func TestMultipleofQuarter(t *testing.T) {
	expectedResults := []int{0, 25}
	// Loop through each receipt object  and compute the total points
	for i, receipt := range receiptsTest {
		expected := expectedResults[i]
		result := multipleofQuarter(receipt.Total)
		if result != expected {
			t.Errorf("Expected %d points, got %d", expected, result)
		}
	}
}
func TestEveryTwoItems(t *testing.T) {
	expectedResults := []int{10, 10}
	// Loop through each receipt object  and compute the total points
	for i, receipt := range receiptsTest {
		expected := expectedResults[i]
		result := everyTwoItems(receipt.Items)
		if result != expected {
			t.Errorf("Expected %d points, got %d", expected, result)
		}
	}
}
func TestTrimmedLength(t *testing.T) {
	expectedResults := []int{6, 0}
	var result int
	var expected int
	// Loop through each receipt object  and compute the total points
	for i, receipt := range receiptsTest {
		expected = expectedResults[i]
		result = trimmedLength(receipt.Items)
	}
	if result != expected {
		t.Errorf("Expected %d points, got %d", expected, result)
	}
}
func TestOddDay(t *testing.T) {
	expectedResults := []int{6, 0}
	// Loop through each receipt object  and compute the total points
	for i, receipt := range receiptsTest {
		expected := expectedResults[i]
		result := oddDay(receipt.PurchaseDate)
		if result != expected {
			t.Errorf("Expected %d points, got %d", expected, result)
		}
	}
}
func TestHappyhour(t *testing.T) {
	expectedResults := []int{0, 10}
	// Loop through each receipt object  and compute the total points
	for i, receipt := range receiptsTest {
		expected := expectedResults[i]
		result := happyhour(receipt.PurchaseTime)
		if result != expected {
			t.Errorf("Expected %d points, got %d", expected, result)
		}
	}
}
