package director

import (
	"foodBuilder/builder"
	"foodBuilder/models"
)

type Director struct {
	builder *builder.PizzaBuilder
}

func NewDirector(builder *builder.PizzaBuilder) *Director {
	return &Director{builder: builder}
}

func (d *Director) ConstructVegetarianPizza() {
	d.builder.SetSize("Large").SetCrust("Thin").SetCheese("Mozzarella").AddTopping("Mushrooms").AddTopping("Bell Peppers").AddTopping("Onions")
}

func (d *Director) ConstructMeatLoversPizza() {
	d.builder.SetSize("Medium").SetCrust("Stuffed").SetCheese("Extra Mozzarella").AddTopping("Pepperoni").AddTopping("Italian Sausage").AddTopping("Ground Beef").AddTopping("Ham")
}

func (d *Director) GetPizza() models.Pizza {
	return d.builder.Build()
}
