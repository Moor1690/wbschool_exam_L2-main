package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

import "fmt"

// Car - продукт, который будет создаваться
type Car struct {
	Make  string
	Color string
	Doors int
}

// CarBuilder - интерфейс строителя
type CarBuilder interface {
	SetMake(make string) CarBuilder
	SetColor(color string) CarBuilder
	SetDoors(doors int) CarBuilder
	Build() Car
}

// carBuilder - конкретный строитель
type carBuilder struct {
	make  string
	color string
	doors int
}

func NewCarBuilder() CarBuilder {
	return &carBuilder{}
}

func (b *carBuilder) SetMake(make string) CarBuilder {
	b.make = make
	return b
}

func (b *carBuilder) SetColor(color string) CarBuilder {
	b.color = color
	return b
}

func (b *carBuilder) SetDoors(doors int) CarBuilder {
	b.doors = doors
	return b
}

func (b *carBuilder) Build() Car {
	return Car{
		Make:  b.make,
		Color: b.color,
		Doors: b.doors,
	}
}

func main() {
	builder := NewCarBuilder()
	car := builder.SetMake("Tesla").SetColor("Red").SetDoors(4).Build()

	fmt.Printf("Car: %+v\n", car)
}
