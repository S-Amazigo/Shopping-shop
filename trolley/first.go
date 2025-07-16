package trolley

// First will be the first thing I upload to github
import (
	"fmt"
	"io"
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
	fmt.Printf("Your total today comes to %d gold coins.\n", sum)
	return sum
}

/*
Tasks for writing API

 0.Extend shop to have metadata, such as item, price, stock level, anything else ()
   Update tests to reflect this change
 1. Function to add to the stock list - shopkeeper would use - will be an API endpoint where we can send data to add to the shop stock
 2. API endpoint to list what we have in stock
	2.1 Another endpoint to get the price of one item in stock
 3. An API endpoint to reduce the stock level of an item
 4. An API endpoint ...

*/

/* Task 0
- Construct an item struct for all the metadata :check:
- Make shop list a slice of above structs
- Rewrite WhatsMyTotal to reference struct entries
*/

// Item is a struct defining what a particular shop item is
type Item struct {
	Price int /* price of item in gold coins per unit quantity */
	// unitQuantity string /* description of unit quantity */ - too much for now?
	StockLevel int  /* number of units of said items in stock  */
	Location   int  /* aisle number product is on*/
	Edible     bool /* Is the item generally edible */
}

// I want the shop stock to be a map of names to item structs to make it easy to add things to and easy to differentiate between an item being out of stock and an item just not being sold by the shop

/*
Note to self about locations:
Aisle 0: Miscallaneous
Aisle 1: Fruits and Veg
Aisle 2: Dairy
Aisle 3: Butcher section
Aisle 4: Kitchen Items
Aisle 5: Home Items
Aisle 6: Baking and baked goods
*/

// shopStock is the list detailing all of the items in the shop
var shopStock = map[string]Item{
	"apple":   {1, 0, 1, true},
	"cabbage": {1, 300, 1, true},
	"milk":    {2, 15, 2, true},
	"cushion": {10, 200, 5, false},
	"salmon":  {5, 0, 3, true},
	"bowl":    {4, 20, 4, false},
}

// WhatsYourTotal given a list of grocery items it outputs the total cost and gold coins and prints a message. I.e. what shop assistant would say at the till
func WhatsYourTotal(w io.Writer, items ...string) int {
	// sum is value of shopping basket
	sum := 0
	for _, v := range items {
		// Make input lowercase in case of capitalisation
		item := strings.ToLower(v)
		check, in := shopStock[item]
		if in {
			if check.StockLevel > 0 {
				sum += check.Price
			} else {
				fmt.Fprintf(w, "Sorry, but we do not have %s in stock \n", item)
			}
		} else {
			fmt.Fprintf(w, "Sorry, but we do not sell %s \n", item)
		}
	}
	fmt.Fprintf(w, "Your total today comes to %d gold coins.\n", sum)
	return sum

}

/* Task 1: Function to add to the stock list - shopkeeper would use - will be an API endpoint where we can send data to add to the shop stock
- Inputs should be name and relevant item information
- Function should create map element and relevant struct
- Set default stockLevel 0 and default location 0 (makes practical sense to me)
- If item is already in stocklist say something like "We stock this already, maybe try AddToStock"
- Shouldn't return anything as we are just updating the stocklist
*/

// AddNewItem adds a new names item to the stock list
func AddNewItem(name string, item Item) {
	if _, haveAlready := shopStock[name]; haveAlready {
		fmt.Printf("We already stock %s. Maybe you want to try ChangeStockBy? \n", name)
	} else {
		shopStock[name] = item
	}
}

/* Task 2: API endpoint to list what we have in stock and add another endpoint to get the price of one item in stock
- Function returns and prints current stocklist
- Function that returns price of stock item
*/

// StockCheck lets you look at the current stock list. It does not return anything. Why? I only want people to be able to make changes to the shopping list through the given commands.
func StockCheck(w io.Writer) {
	fmt.Fprintf(w, "%v, \n", shopStock)
}

func CheckPrice(w io.Writer, name string) int {
	if item, haveAlready := shopStock[name]; haveAlready {
		return item.Price
	} else {
		fmt.Fprintf(w, "We don't stock %s. \n", name)
		return 0
	}
}

/* Task 3: An API endpoint to reduce the stock level of an item
- Check numbers work as expected
- Print relevant message for unstocked items
*/
// ChangeStockBy changes the stock level of a given item by a given amount
func ChangeStockBy(name string, amount int) {
	if item, haveAlready := shopStock[name]; haveAlready {
		item.StockLevel += amount
	} else {
		fmt.Printf("We don't stock %s. Maybe you want to try AddNewItem? \n", name)
	}
}

// Command just reverts stock list to original, as defined in package
func ResetStock() {
	point2Stock := &shopStock
	*point2Stock = map[string]Item{
		"apple":   {1, 0, 1, true},
		"cabbage": {1, 300, 1, true},
		"milk":    {2, 15, 2, true},
		"cushion": {10, 200, 5, false},
		"salmon":  {5, 0, 3, true},
		"bowl":    {4, 20, 4, false},
	}
}

/* Possible other stuff
- Store markup/markdown raise all prices by a set amount
- Delete items
- Selling - input list of purchases and how much, output total cost and change stocklevels
- Set stock levels
- Set prices
- Set stock to default (as explicitly defined in package)
*/
