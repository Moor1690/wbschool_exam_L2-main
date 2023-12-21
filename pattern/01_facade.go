package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/
import "fmt"

/*
   Паттерн "Фасад" упрощает сложные системы, предоставляя простой интерфейс.
   В этом примере фасад используется для управления различными компонентами компьютера.
*/

func main() {
	computer := NewComputer()
	computer.Start()
	computer.Shutdown()
}

// RAM управляет оперативной памятью
type RAM struct{}

func (r *RAM) Load(position int, data string) {
	fmt.Printf("Loading data '%s' at position %d\n", data, position)
}

func (r *RAM) Free(position int) {
	fmt.Printf("Freeing memory at position %d\n", position)
}

// CPU управляет процессором
type CPU struct{}

func (c *CPU) Execute() {
	fmt.Println("Executing commands")
}

func (c *CPU) Halt() {
	fmt.Println("Halting execution")
}

// HardDrive управляет жестким диском
type HardDrive struct{}

func (hd *HardDrive) Read(position int, size int) string {
	return fmt.Sprintf("Reading %d bytes from position %d", size, position)
}

func (hd *HardDrive) Write(position int, data string) {
	fmt.Printf("Writing data '%s' at position %d\n", data, position)
}

// Computer (Facade) упрощает управление компонентами компьютера
type Computer struct {
	ram       RAM
	cpu       CPU
	hardDrive HardDrive
}

func NewComputer() *Computer {
	return &Computer{
		ram:       RAM{},
		cpu:       CPU{},
		hardDrive: HardDrive{},
	}
}

func (c *Computer) Start() {
	c.ram.Load(0, "OS Data")
	c.cpu.Execute()
	c.hardDrive.Read(0, 1024)
}

func (c *Computer) Shutdown() {
	c.cpu.Halt()
	c.ram.Free(0)
}
