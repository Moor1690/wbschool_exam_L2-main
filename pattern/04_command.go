package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/
import "fmt"

// Интерфейс для команд
type Command interface {
	Execute()
}

// Конкретные команды

type LightOnCommand struct {
	Light *Light
}

func (c *LightOnCommand) Execute() {
	c.Light.On()
}

type LightOffCommand struct {
	Light *Light
}

func (c *LightOffCommand) Execute() {
	c.Light.Off()
}

// Получатель команды
type Light struct{}

func (l *Light) On() {
	fmt.Println("Свет включен")
}

func (l *Light) Off() {
	fmt.Println("Свет выключен")
}

// Инвокер (отправитель команды)
type RemoteControl struct {
	Command Command
}

func (rc *RemoteControl) PressButton() {
	rc.Command.Execute()
}

func main() {
	light := &Light{}
	lightOnCommand := &LightOnCommand{Light: light}
	lightOffCommand := &LightOffCommand{Light: light}

	remote := &RemoteControl{}

	// Включение света
	remote.Command = lightOnCommand
	remote.PressButton()

	// Выключение света
	remote.Command = lightOffCommand
	remote.PressButton()
}
