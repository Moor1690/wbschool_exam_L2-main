package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/
import "fmt"

// Интерфейс для обработчика
type Handler interface {
	Handle(request string) bool
	SetNext(handler Handler)
}

// Базовая реализация обработчика
type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handler Handler) {
	b.next = handler
}

// Конкретные обработчики

type ConcreteHandlerA struct{}

func (c *ConcreteHandlerA) Handle(request string) bool {
	if request == "A" {
		fmt.Println("Обработчик A обработал запрос")
		return true
	} else if c.next != nil {
		return c.next.Handle(request)
	}
	return false
}

type ConcreteHandlerB struct{}

func (c *ConcreteHandlerB) Handle(request string) bool {
	if request == "B" {
		fmt.Println("Обработчик B обработал запрос")
		return true
	} else if c.next != nil {
		return c.next.Handle(request)
	}
	return false
}

type ConcreteHandlerC struct{}

func (c *ConcreteHandlerC) Handle(request string) bool {
	if request == "C" {
		fmt.Println("Обработчик C обработал запрос")
		return true
	} else {
		fmt.Println("Ни один обработчик не может обработать запрос")
		return false
	}
}

func main() {
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}
	handlerC := &ConcreteHandlerC{}

	handlerA.SetNext(handlerB)
	handlerB.SetNext(handlerC)

	requests := []string{"A", "B", "C", "D"}

	for _, request := range requests {
		if !handlerA.Handle(request) {
			fmt.Printf("Запрос %s не может быть обработан\n", request)
		}
	}
}
