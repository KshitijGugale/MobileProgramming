package main

import (
	"fmt"
	"os"
)

type Data struct {
	ID int
	SKU	string
	Name string
	Price float32
	Qty int
}

var Items = []Data {
	Data {
		ID: 1,
		SKU: "120P90",
		Name: "Google Home",
		Price: 49.99,
		Qty: 10,
	},
	Data {
		ID: 2,
		SKU: "43N23P",
		Name: "MacBook Pro",
		Price: 5399.99,
		Qty: 5,
	},
	Data {
		ID: 3,
		SKU: "A304SD",
		Name: "Alexa Speaker",
		Price: 109.50,
		Qty: 10,
	},
	Data {
		ID: 4,
		SKU: "234234",
		Name: "Raspberry Pi B",
		Price: 30,
		Qty: 2,
	},
}
func main () {
	PrintDetails()

	var id, qty int
	fmt.Println("What you want to buy? Choose ID from 1 to 4.")
	fmt.Scanf("%d", &id)
	if id < 1 || id >= 5 {
		fmt.Println("Invalid ID choosen!")
		os.Exit(0);
	}
	fmt.Printf("How many items of %s you want?\n", Items[id-1].Name)
	fmt.Scanf("%d", &qty)
	price := CalculatePrice(id, qty)
	fmt.Printf("Total Price: %0.2f\n", price)
}

func CalculatePrice(id, qty int) float32 {
	var totalPrice float32
	if id == 2 { //MacBook Pro
		var div int = qty / 3
		var rem int = qty % 3

		totalPrice += Items[id-1].Price * float32(div) * 2
		totalPrice += Items[id-1].Price * float32(rem)
	} else if id == 3 { //Alexa Speaker
		totalPrice = totalPrice + Items[id-1].Price * float32(qty)
		if qty > 3 {
			totalPrice -= (0.1) * totalPrice
		}
	} else {
		totalPrice = totalPrice + Items[id-1].Price * float32(qty)
	}

	return totalPrice
}

func PrintDetails() {
	for _, value := range Items {
		fmt.Printf("ID: %d, SKU: %s, Name: %s, Price: %f, Qty: %d\n", value.ID, value.SKU, value.Name, value.Price, value.Qty)
	}
}