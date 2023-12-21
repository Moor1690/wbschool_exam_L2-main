package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/
import "fmt"

// Интерфейс для посетителя
type Visitor interface {
	VisitCircle(circle *Circle)
	VisitRectangle(rectangle *Rectangle)
	VisitTriangle(triangle *Triangle)
}

// Абстрактный тип геометрической фигуры
type Shape interface {
	Accept(visitor Visitor)
}

// Конкретные типы фигур

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

type Triangle struct {
	Side1, Side2, Side3 float64
}

// Реализация методов Accept для каждой фигуры
func (c *Circle) Accept(visitor Visitor) {
	visitor.VisitCircle(c)
}

func (r *Rectangle) Accept(visitor Visitor) {
	visitor.VisitRectangle(r)
}

func (t *Triangle) Accept(visitor Visitor) {
	visitor.VisitTriangle(t)
}

// Конкретный посетитель для вычисления площади и периметра фигур
type AreaPerimeterVisitor struct {
	Area      float64
	Perimeter float64
}

func (v *AreaPerimeterVisitor) VisitCircle(circle *Circle) {
	v.Area += 3.14159265 * circle.Radius * circle.Radius
	v.Perimeter += 2 * 3.14159265 * circle.Radius
}

func (v *AreaPerimeterVisitor) VisitRectangle(rectangle *Rectangle) {
	v.Area += rectangle.Width * rectangle.Height
	v.Perimeter += 2 * (rectangle.Width + rectangle.Height)
}

func (v *AreaPerimeterVisitor) VisitTriangle(triangle *Triangle) {
	// Пусть здесь будет простое приближение, не учитывая форму треугольника
	v.Area += 0.5 * triangle.Side1 * triangle.Side2
	v.Perimeter += triangle.Side1 + triangle.Side2 + triangle.Side3
}

func main() {
	circle := &Circle{Radius: 5}
	rectangle := &Rectangle{Width: 4, Height: 6}
	triangle := &Triangle{Side1: 3, Side2: 4, Side3: 5}

	visitor := &AreaPerimeterVisitor{}

	// Вычисляем площадь и периметр каждой фигуры
	circle.Accept(visitor)
	rectangle.Accept(visitor)
	triangle.Accept(visitor)

	fmt.Printf("Площадь всех фигур: %.2f\n", visitor.Area)
	fmt.Printf("Периметр всех фигур: %.2f\n", visitor.Perimeter)
}
