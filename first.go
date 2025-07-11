// Practice is where my additional practice exercises live i.e. those not covered in Tour of Go or Exercism
package main

// First will be the first thing I upload to github
import (
	"fmt"
	"strings"
)

// shopPrices is a map of shop items and their unit price in gold coins
var shopPrices = map[string]int{
	"fruit":     2,
	"vegetable": 2,
	"meat":      5,
	"milk":      3,
	"bowl":      4,
	"cutlery":   4,
}

// WhatsMyTotal given a list of grocery items it outputs the total cost and gold coins and prints a message
func WhatsMyTotal(items ...string) int {
	// sum is value of shopping basket
	sum := 0
	for _, v := range items {
		// Make input lowercase in case of capitalisation
		item := strings.ToLower(v)
		check, in := shopPrices[item]
		if in {
			sum += check
		} else {
			fmt.Printf("Sorry, but we do not have %s in stock \n", item)
		}
	}
	fmt.Printf("Your total today comes to %d gold coins.", sum)
	return sum
}

func main() {
	WhatsMyTotal("fruit", "vegetable", "milk", "bowl")
}
