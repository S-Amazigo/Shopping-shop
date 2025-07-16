package trolley

import (
	"bytes"
	"testing"
)

func TestWhatsMyTotal(t *testing.T) {

	/*
		Test Driven Development (TDD) - like exercism, best practice

		// set up the test with a case
		// define the expected outcome
		// compare the actual outcome with the expected outcome
		// if something in the list is not present, expect an output of "sorry ..."

		2 cases:
		1. When the list contains some items, it should return the price.
			If there is a mixture of upper and lower case, the function shouldn't care about it
		2. When the list is missing an item, it should return "sorry ...".
	*/
	t.Run("When the list contains items in the shop, it should return the total price of the items", func(t *testing.T) {
		/*
			1. Add milk, meat and a bowl, total should be 12 gold coins
				Expected return value: 12
				Print out "Your total today comes to 12 gold coins."
		*/

		got := WhatsMyTotal("milk", "meat", "bowl")
		expected := 12
		if got != expected {
			t.Fatal("expected", expected, "but got", got)
		}
		t.Log("Success!")
	})

	/*
		For the second test, should think of cases where the function wouldn't have desired output if run the function didn't do what it should
	*/

	t.Run("The function should interpret all inputs as if there were lowercase ", func(t *testing.T) {
		// Try some crazy capitalisations and functions should still return as normal
		got := WhatsMyTotal("Milk", "mEAt", "BOWL")
		expected := 12
		if got != expected {
			t.Fatal("expected", expected, "but got", got)
		}
		t.Log("Success!")

	})

	//Check print message - How do I do that?
	t.Run("When given an item is not on the shopPrice list, relevant apology should be printed", func(t *testing.T) {
		got := "a"
		expected := "Sorry, but we do not have eggs in stock \n"
		if got != expected {
			t.Fatal("expected", expected, "but got", got)
		}
		t.Log("Success!")
	})
}

func TestWhatsYourTotal(t *testing.T) {
	/* Cases:
	   1. If all given items are on the stocklist and are in stock the function should return the total value and print message communicating the total
	   2. If all given items are on the stocklist but some are out of stock it should print the relavent message
	   3. If some given items are not on the list function should print out message saying that we don't sell that
	*/

	// Case 1 - Input in stock
	t.Run("If all given items are on the stocklist and are in stock the function should return the total value and print message communicating the total", func(t *testing.T) {
		var outputMessage bytes.Buffer
		gotSum := WhatsYourTotal(&outputMessage, "cushion", "cabbage", "bowl")
		expectedSum := 15
		expectedMessage := "Your total today comes to 15 gold coins.\n"

		if outputMessage.String() != expectedMessage {
			t.Fatal("Wrong message!!! expected:", expectedMessage, "but got:", outputMessage.String())
		} else if gotSum != expectedSum {
			t.Fatal("Wrong sum!!! \n expected:", expectedSum, "but got:", gotSum)
		}
		t.Log("Success!")
	})

	// Case 2 - Input items exist but are not in stock
	t.Run("If all given items are on the stocklist but some are out of stock it should print the relavent message", func(t *testing.T) {
		var outputMessage bytes.Buffer
		gotSum := WhatsYourTotal(&outputMessage, "apple", "salmon")
		expectedSum := 0
		expectedMessage := "Sorry, but we do not have apple in stock \nSorry, but we do not have salmon in stock \nYour total today comes to 0 gold coins.\n"

		if outputMessage.String() != expectedMessage {
			t.Fatal("Wrong message!!! expected:", expectedMessage, "but got:", outputMessage.String())
		} else if gotSum != expectedSum {
			t.Fatal("Wrong sum!!! \n expected:", expectedSum, "but got:", gotSum)
		}
		t.Log("Success!")
	})

	// Case 3 - Input items don't exist
	t.Run("If some given items are not on the list function should print out message saying that we don't sell that", func(t *testing.T) {
		var outputMessage bytes.Buffer
		gotSum := WhatsYourTotal(&outputMessage, "beans")
		expectedSum := 0
		expectedMessage := "Sorry, but we do not sell beans \nYour total today comes to 0 gold coins.\n"

		if outputMessage.String() != expectedMessage {
			t.Fatal("Wrong message!!! expected:", expectedMessage, "but got:", outputMessage.String())
		} else if gotSum != expectedSum {
			t.Fatal("Wrong sum!!! \n expected:", expectedSum, "but got:", gotSum)
		}
		t.Log("Success!")
	})
}

func TestStockCheck(t *testing.T) {
	// Check map prints as expected and map is returned

}

func TestAddNewItem(t *testing.T) {
	/*

	   	Cases:

	   1. Check empty case
	   2. If no preexisting items are given map extends appropriately
	   3. If preexisting item is given appropriate message is printed
	*/
}

func TestCheckPrice(t *testing.T) {
	/* Cases:
	   1. If non existent item asked for print appropriate message
	   2. Display corrrect prices for existing items
	*/
}

func TestChangeStockBy(t *testing.T) {
	/* Cases:
	   1. No stock changes
	   2. Positive and negative stock changes
	   3. Print relevant items if not in stock
	*/

}

/*  Example demonstrating interfaces
type MyWriter struct{}

func (m MyWriter) Jess() bool {
	return true
}

func (m MyWriter) Stephanie() string {
	return "Stephanie"
}

func (m MyWriter) Write(p []byte) (n int, err error) {
	fmt.Println("I don't care what you've written")
	return 0, nil
}
*/
