package builder

import "foodBuilder/models"

type PizzaBuilder struct {
	pizza models.Pizza
}

func NewPizzaBuilder() *PizzaBuilder {
	return &PizzaBuilder{pizza: models.Pizza{}}
}

func (pb *PizzaBuilder) SetSize(size string) *PizzaBuilder {
	pb.pizza.Size = size

	return pb
}

func (pb *PizzaBuilder) SetCrust(crust string) *PizzaBuilder {
	pb.pizza.Crust = crust

	return pb
}

func (pb *PizzaBuilder) SetCheese(cheese string) *PizzaBuilder {
	pb.pizza.Cheese = cheese

	return pb
}

func (pb *PizzaBuilder) AddTopping(topping string) *PizzaBuilder {
	pb.pizza.Toppings = append(pb.pizza.Toppings, topping)

	return pb
}

func (pb *PizzaBuilder) Build() models.Pizza {
	return pb.pizza
}
