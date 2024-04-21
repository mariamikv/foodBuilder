package main

import (
	"fmt"
	"foodBuilder/builder"
	"foodBuilder/director"
)

func main() {

	pizzaBuilder := builder.NewPizzaBuilder()
	pizzaDirector := director.NewDirector(pizzaBuilder)

	pizzaDirector.ConstructVegetarianPizza()
	vegetarianPizza := pizzaDirector.GetPizza()
	fmt.Printf("Vegetarian Pizza details:\n Size: %s\n Crust: %s\n Cheese: %s\n Toppings: %v\n", vegetarianPizza.Size, vegetarianPizza.Crust, vegetarianPizza.Cheese, vegetarianPizza.Toppings)

	pizzaDirector.ConstructMeatLoversPizza()
	meatLoversPizza := pizzaDirector.GetPizza()
	fmt.Printf("Meat Lover's Pizza details:\n Size: %s\n Crust: %s\n Cheese: %s\n Toppings: %v\n", meatLoversPizza.Size, meatLoversPizza.Crust, meatLoversPizza.Cheese, meatLoversPizza.Toppings)
}
