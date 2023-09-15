package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Burger struct {
	Bun     bool
	Price   float64
	Dressed bool
}

type Drink struct {
	Price float64
	Type  string
}

type Side struct {
	Price float64
	Type  string
}

type Combo struct {
	burger Burger
	side   Side
	drink  Drink
	price  float64
}

func user_input_burger() Burger {
	var inputBurger Burger
	var inputBun string
	var inputDress string
	fmt.Print("Do you want Burger with Bun?")
	fmt.Scan(&inputBun)
	if inputBun == "Yes" || inputBun == "yes" {
		inputBurger.Bun = true
	} else if inputBun == "No" || inputBun == "no" {
		inputBurger.Bun = false
	}
	fmt.Print("Do you want Burger dressed?")
	fmt.Scan(&inputDress)
	if inputDress == "Yes" || inputDress == "yes" {
		inputBurger.Dressed = true
	} else if inputDress == "No" || inputDress == "no" {
		inputBurger.Dressed = false
	}
	if inputBurger.Bun {
		inputBurger.Price = 7.0
	} else {
		inputBurger.Price = 6.0
	}
	return inputBurger
}

func user_input_drink() Drink {
	var inputDrink Drink
	fmt.Print("Do you want water or soda?")
	var inputDrinkType string
	fmt.Scan(&inputDrinkType)
	if inputDrinkType == "water" {
		inputDrink.Price = 1.0
		inputDrink.Type = "water"
	} else if inputDrinkType == "soda" {
		inputDrink.Price = 2.0
		inputDrink.Type = "soda"
	}
	return inputDrink
}

func user_input_side() Side {
	var inputSide Side
	var inputSideOption string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want side as fries, onion rings, a salad, or coleslaw ?")
	inputSideOption, _ = reader.ReadString('\n')
	inputSideOption = strings.Replace(inputSideOption, "\n", "", -1)
	fmt.Println(inputSideOption)
	if inputSideOption == "fries" {
		inputSide.Price = 2.0
		inputSide.Type = "fries"
	} else {
		fmt.Println(inputSideOption)
		inputSide.Price = 3.0
		inputSide.Type = inputSideOption
	}

	return inputSide
}

func userInputCombo() Combo {
	burger := user_input_burger()
	var inputCombo Combo
	var inputComboOption string
	fmt.Print("Do you want burger with side or drink?")
	fmt.Scan(&inputComboOption)
	if inputComboOption == "side" {
		inputCombo.side = user_input_side()
	} else if inputComboOption == "drink" {
		inputCombo.drink = user_input_drink()
	}
	inputCombo.burger = burger
	return inputCombo
}

func orderDetails(combo Combo) float64 {
	comboPrice := combo.price + combo.burger.Price

	if combo.side.Type != "" {
		comboPrice += combo.side.Price
	}
	if combo.drink.Type != "" {
		comboPrice += combo.drink.Price
	}

	fmt.Println("Order details: ")
	fmt.Printf("Burger: Bun - %v, Dressed - %v, Price - $%.2f\n",
		combo.burger.Bun, combo.burger.Dressed, combo.burger.Price)
	if combo.burger.Bun {
		fmt.Printf("Combo Burger: Bun - %v, Dressed - %v, Price - $%.2f\n",
			combo.burger.Bun, combo.burger.Dressed, combo.burger.Price)
	}
	if combo.side.Type != "" {
		fmt.Printf("Side: Type - %v, Price - $%.2f\n", combo.side.Type, combo.side.Price)
	}
	if combo.drink.Type != "" {
		fmt.Printf("Drink: Type - %v, Price - $%.2f\n", combo.drink.Type, combo.drink.Price)
	}
	fmt.Printf("Total price for combo: $%.2f\n", comboPrice)
	return comboPrice

}

func takeOrderFromUser() {
	fmt.Print("Name of your order: ")
	var orderName string
	fmt.Scan(&orderName)
	var done string
	var totalPrice float64

	for done != "done" {
		fmt.Print("Please type 'y' to continue ordering or 'done' to finish: ")
		fmt.Scan(&done)
		if done != "done" {
			combo := userInputCombo()
			orderPrice := orderDetails(combo)
			totalPrice = totalPrice + orderPrice
		} else {
			fmt.Println("Thank you for your order!")
			fmt.Print("Total price for your order is:", totalPrice)
		}
	}

	cat ~/.ssh/id_rsa.pub.ssh/config

	git remote set-url origin https://YOUR_TOKEN@github.com/username/repo.git}

func main() {
	fmt.Println("Welcome to our burger shop!")
	takeOrderFromUser()
}
