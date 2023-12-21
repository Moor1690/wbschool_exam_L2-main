package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/
import "fmt"

// Интерфейс для фабрики
type Factory interface {
	CreateProduct() Product
}

// Интерфейс для продукта
type Product interface {
	Use()
}

// Конкретные продукты

type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() {
	fmt.Println("Продукт A используется")
}

type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() {
	fmt.Println("Продукт B используется")
}

// Конкретные фабрики

type ConcreteFactoryA struct{}

func (f *ConcreteFactoryA) CreateProduct() Product {
	return &ConcreteProductA{}
}

type ConcreteFactoryB struct{}

func (f *ConcreteFactoryB) CreateProduct() Product {
	return &ConcreteProductB{}
}

func main() {
	// Клиентский код
	factoryA := &ConcreteFactoryA{}
	productA := factoryA.CreateProduct()
	productA.Use()

	factoryB := &ConcreteFactoryB{}
	productB := factoryB.CreateProduct()
	productB.Use()
}
