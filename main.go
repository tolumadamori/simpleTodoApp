package main

import (
	"fmt"
	"os"
)

type todoItem struct {
	id   int
	todo string
}

var listOfItems []todoItem

func menu() {
	fmt.Println("Please select an option:........\n #1 - Create an item \t \t #2 - List all items \n #3 - Delete an item \t \t #4 - Update an item \n \t #5 - Exit program")
	fmt.Println("\n enter the option number and press enter:")
	var option int

	fmt.Scanln(&option)

	switch option {
	case 1:
		createItems()
	case 2:
		listItems()
	case 3:
		deleteItems()
		menu()

	case 4:
		updateItems()
		// menu()

	case 5:
		exit()

	default:
		fmt.Println("The option you selected is invalid. Please select a valid option")
		menu()
	}

}

func createItems() {

	fmt.Println("Type in the new item and press enter")
	var newItem todoItem
	fmt.Scanln(&newItem.todo)
	newItem.id = len(listOfItems) + 1
	listOfItems = append(listOfItems, newItem)
	menu()

}

func listItems() {
	fmt.Println("\n All your todo items are listed below: \n")
	for _, v := range listOfItems {
		fmt.Printf("Item %v is %v \n", v.id, v.todo)
	}

	fmt.Println("\n \n")
	menu()
}

func deleteItems() {
	fmt.Println("Enter the id of the item you want to delete or enter '0' to go back")
	var itemToDelete int
	fmt.Scanln(&itemToDelete)
	if itemToDelete == 0 {
		menu()
	} else if itemToDelete < 0 || itemToDelete > len(listOfItems)+1 {
		fmt.Println("Invalid item id entered")
		deleteItems()
	} else {
		for i, v := range listOfItems {
			if v.id == itemToDelete {
				fmt.Printf("\n \n item %v - %v is deleted \n \n", v.id, v.todo)
				listOfItems = append(listOfItems[:i], listOfItems[i+1:]...)

			}
		}

	}
}

func updateItems() {
	fmt.Println("Enter the id of the item you want to Update or enter '0' to go back")
	var itemToUpdate int
	fmt.Scanln(&itemToUpdate)

	if itemToUpdate == 0 {
		menu()
	} else if itemToUpdate < 0 || itemToUpdate > len(listOfItems) {
		fmt.Println("Invalid item id entered")
		updateItems()
	} else {

		for i, v := range listOfItems {
			if v.id == itemToUpdate {
				listOfItems = append(listOfItems[:i], listOfItems[i+1:]...)
				fmt.Println("enter the new item")
				var updatedItem todoItem
				updatedItem.id = v.id
				fmt.Scanln(&updatedItem.todo)
				listOfItems = append(listOfItems, updatedItem)
				fmt.Printf("item %v -  is now updated to %v \n \n", updatedItem.id, updatedItem.todo)

			}
		}
		menu()

	}

}

func exit() {
	os.Exit(1)

}
func main() {
	fmt.Println("\n \n This is the start of the program")
	menu()
}
