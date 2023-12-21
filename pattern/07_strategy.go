package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/
import "fmt"

// Интерфейс для стратегии
type PaymentStrategy interface {
	Pay(amount float64)
}

// Конкретные стратегии

type CreditCardPayment struct{}

func (c *CreditCardPayment) Pay(amount float64) {
	fmt.Printf("Оплачено %.2f с помощью кредитной карты\n", amount)
}

type PayPalPayment struct{}

func (p *PayPalPayment) Pay(amount float64) {
	fmt.Printf("Оплачено %.2f через PayPal\n", amount)
}

// Контекст (клиентский код)

type ShoppingCart struct {
	PaymentMethod PaymentStrategy
}

func (cart *ShoppingCart) Checkout(amount float64) {
	cart.PaymentMethod.Pay(amount)
}

func main() {
	creditCardPayment := &CreditCardPayment{}
	payPalPayment := &PayPalPayment{}

	cart := &ShoppingCart{}

	// Оплата с помощью кредитной карты
	cart.PaymentMethod = creditCardPayment
	cart.Checkout(100.00)

	// Оплата через PayPal
	cart.PaymentMethod = payPalPayment
	cart.Checkout(50.00)
}
