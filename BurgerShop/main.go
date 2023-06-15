package main

import (
	"fmt"
	"strings"
)

const burgerPrice = 6.00
const sidePrice = 2.00
const comboDiscount = 1.00

var burgerCondiments = []string{"Lettuce", "Tomato", "Onion", "Mayo"}
var drinkTypes = []string{"FANTA", "COKE", "SPRITE", "PEPSI"}
var drinks = map[int]int{12: 1, 16: 2, 24: 3}
var sideTypes = []string{"fries", "coleslaw", "salad"}
var possibleChoices = []string{"b", "s", "d", "c"}

type burger struct {
	name       string
	price      int
	condiments []string
}

func (b *burger) getName() string {
	return b.name
}
func (b *burger) computePrice() int {
	b.price = burgerPrice
	return b.price
}
func (b *burger) addCondiment(condiment string) {
	b.condiments = append(b.condiments, condiment)
}

func (b *burger) display(displayPrice bool) {
	fmt.Println("Item Name: " + b.getName())
	fmt.Print("Condiments: ")
	for _, condiment := range b.condiments {
		fmt.Print(condiment + " ")
	}
	fmt.Println()
	if displayPrice == true {
		fmt.Printf("Item Price: $%d\n", b.computePrice())
	}

}

type drink struct {
	name  string
	size  int
	price int
}

func (d *drink) getName() string {
	return d.name
}
func (d *drink) getSize() int {
	return d.size
}

func (d *drink) computePrice() int {
	if _, ok := drinks[d.getSize()]; ok {
		d.price = drinks[d.getSize()]
	}
	return d.price
}

func (d *drink) display(displayPrice bool) {
	fmt.Println("Item Name: " + strings.ToUpper(d.getName()))
	fmt.Printf("Item Size: %d\n", d.size)
	if displayPrice == true {
		fmt.Printf("Item Price: $%d\n", d.computePrice())
	}
}

type side struct {
	name  string
	price int
}

func (s *side) getName() string {
	return s.name
}
func (s *side) computePrice() int {
	s.price = sidePrice
	return s.price
}
func (s *side) display(displayPrice bool) {

	fmt.Println("Item Name: " + s.getName())
	if displayPrice == true {
		fmt.Printf("Item Price: $%d\n", s.computePrice())
	}
}

type combo struct {
	name   string
	burger burger
	drink  drink
	side   side
	price  int
}

func (c *combo) getName() string {
	return c.name
}

func (c *combo) computePrice() int {
	c.price = c.burger.computePrice() + c.drink.computePrice() + c.side.computePrice() - comboDiscount
	return c.price
}

func (c *combo) display() {
	fmt.Println("Burger For Combo")
	c.burger.display(false)
	fmt.Println("Side For Combo")
	c.side.display(false)
	fmt.Println("Drink For Combo")
	c.drink.display(false)
	fmt.Printf("Price For Combo: $%d\n", c.computePrice())
}

type order struct {
	name    string
	price   int
	burgers []burger
	drinks  []drink
	sides   []side
	combos  []combo
}

func (o *order) getName() string {
	return o.name
}

func (o *order) computePrice() int {
	var price = 0

	for _, b := range o.burgers {
		price = price + b.computePrice()
	}
	for _, s := range o.sides {
		price = price + s.computePrice()
	}
	for _, d := range o.drinks {
		price = price + d.computePrice()
	}
	for _, c := range o.combos {
		price = price + c.computePrice()

	}
	o.price = price
	return o.price
}

func (o *order) display() {
	fmt.Println("====================================")
	fmt.Println("===========ORDER OVERVIEW===========")
	for k, b := range o.burgers {
		fmt.Printf("=====Burger %d\n", k+1)
		b.display(true)
	}
	for k, s := range o.sides {
		fmt.Printf("=====Side %d\n", k+1)
		s.display(true)
	}
	for k, d := range o.drinks {
		fmt.Printf("=====Drink %d\n", k+1)
		d.display(true)
	}
	for k, c := range o.combos {
		fmt.Printf("=====Combo %d\n", k+1)
		c.display()
	}
	fmt.Printf("=====ORDER TOTAL: $%d\n", o.computePrice())
	fmt.Println("====================================")

}

func contains(arr []string, choice string) bool {
	for _, v := range arr {
		if v == choice {
			return true
		}
	}
	return false
}
func orderBurger() burger {
	var b burger
	b.name = "Beef Burger"
	fmt.Print("Do you want condiments on your burger? (type y for yes): ")
	var choice1 string
	fmt.Scanln(&choice1)
	if strings.ToLower(choice1) == "y" {
		for _, condiment := range burgerCondiments {
			var choice2 string
			fmt.Print("Do you want " + condiment + " on your burger? (type y for yes): ")
			fmt.Scanln(&choice2)
			if strings.ToLower(choice2) == "y" {
				b.addCondiment(condiment)

			}
		}

	}
	return b

}
func orderSide() side {
	fmt.Print("These are the available sides: ")
	fmt.Println(sideTypes)
	var choice bool = false
	var sideTypeChoice string
	for choice == false {
		fmt.Print("What side do you want? ")
		fmt.Scanln(&sideTypeChoice)
		if contains(sideTypes[:], sideTypeChoice) {
			choice = true
		} else {
			fmt.Println("Please enter a valid choice")
		}
	}
	var s side
	s.name = strings.ToLower(sideTypeChoice)
	s.computePrice()
	return s
}

func orderDrink() drink {
	fmt.Print("These are the available drinks: ")
	fmt.Println(drinkTypes)
	fmt.Print("These are the available sizes: ")
	fmt.Println("[12 16 24]")

	var choice bool = false
	var drinkTypeChoice string
	var drinkSizeChoice int
	for choice == false {
		fmt.Print("What drink do you want? ")
		fmt.Scanln(&drinkTypeChoice)
		if contains(drinkTypes, strings.ToUpper(drinkTypeChoice)) {
			choice = true
		} else {
			fmt.Println("Please enter a valid drink")
		}
	}
	choice = false
	for choice == false {
		fmt.Print("What size do you want? ")
		fmt.Scanln(&drinkSizeChoice)
		if _, ok := drinks[drinkSizeChoice]; ok {
			choice = true
		} else {
			fmt.Println("Please enter a valid size")
		}
	}
	var d drink
	d.name = strings.ToLower(drinkTypeChoice)
	d.size = drinkSizeChoice
	d.computePrice() // equivalent also to d.price = drinks[drinkSizeChoice]
	return d
}

func orderCombo() combo {
	var c combo
	fmt.Println("Let's get you a combo meal!")
	fmt.Println("First, let's order the burger for your combo")
	c.burger = orderBurger()

	fmt.Println("Now, let's order the drink for your combo")
	c.drink = orderDrink()

	fmt.Println("Finally, let's order the side for your combo")
	c.side = orderSide()

	return c
}

func main() {
	var ord order
	var name string
	var done bool
	done = false
	fmt.Println("Welcome to Myriam's Burger Shop!")
	fmt.Print("May I have your name for the order?  ")
	fmt.Scanln(&name)
	ord.name = name
	fmt.Println("Let's get your order in " + name + "!")
	for done == false {
		fmt.Println("Enter b for Burger")
		fmt.Println("Enter s for Side")
		fmt.Println("Enter d for Drink")
		fmt.Print("Enter c for Combo: ")
		choice := ""
		for contains(possibleChoices[:], choice) == false {
			fmt.Scanln(&choice)
			switch choice {
			case "b":
				fmt.Println("Burger it is!")
				var b = orderBurger()
				ord.burgers = append(ord.burgers, b)
			case "s":
				fmt.Println("Side it is!")
				var s = orderSide()
				ord.sides = append(ord.sides, s)
			case "d":
				fmt.Println("Drink it is!")
				var d = orderDrink()
				ord.drinks = append(ord.drinks, d)
			case "c":
				fmt.Println("Combo it is!")
				var c = orderCombo()
				ord.combos = append(ord.combos, c)
			default:
				fmt.Println("Unknown choice")
				fmt.Println("Please enter a valid choice")
			}
		}
		fmt.Print("Do you want to order more items? (Enter n or N to stop.):   ")
		var q1 string
		fmt.Scanln(&q1)
		if strings.ToLower(q1) == "n" {
			done = true
		}
	}
	ord.display()
}
