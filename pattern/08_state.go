package pattern

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/
import "fmt"

// Интерфейс для состояния
type state interface {
	getMeal(*dayMeals)
}

// Структура, представляющая меню на день
type dayMeals struct {
	s state
}

func (dm *dayMeals) setState(s state) {
	dm.s = s
}

func (dm *dayMeals) request() {
	dm.s.getMeal(dm)
}

// Конкретное состояние: Завтрак
type breakfast struct{}

func (b *breakfast) getMeal(dm *dayMeals) {
	fmt.Println("Завтрак: Вот ваш омлет и жареный бекон на завтрак!")
	dm.setState(&dinner{})
}

// Конкретное состояние: Обед
type dinner struct{}

func (d *dinner) getMeal(dm *dayMeals) {
	fmt.Println("Обед :Ваш великолепный борщ с зеленым луком уже на столе!")
	dm.setState(&supper{})
}

// Конкретное состояние: Ужин
type supper struct{}

func (s *supper) getMeal(dm *dayMeals) {
	fmt.Println("Ужин: Бефстроганов с жареным картофелем остывает. Поторопись!")
	dm.setState(&breakfast{})
}

func main() {
	// Создаем экземпляр структуры, представляющей меню на день
	dm := &dayMeals{s: &breakfast{}}

	// Симулируем запросы на прием пищи
	dm.request()
	dm.request()
	dm.request()
	dm.request()
}
