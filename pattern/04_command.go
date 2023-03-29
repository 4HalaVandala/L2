package main

import "fmt"

// Receiver interface
type Light interface {
	On()
	Off()
}

// Concrete receiver
type BedroomLight struct{}

func (l *BedroomLight) On() {
	fmt.Println("Bedroom light is on")
}

func (l *BedroomLight) Off() {
	fmt.Println("Bedroom light is off")
}

// Command interface
type Command interface {
	Execute()
}

// Concrete command
type TurnOnLightCommand struct {
	Light Light
}

func (c *TurnOnLightCommand) Execute() {
	c.Light.On()
}

// Concrete command
type TurnOffLightCommand struct {
	Light Light
}

func (c *TurnOffLightCommand) Execute() {
	c.Light.Off()
}

// Invoker
type Switch struct {
	OnCommand  Command
	OffCommand Command
}

func (s *Switch) FlipOn() {
	s.OnCommand.Execute()
}

func (s *Switch) FlipOff() {
	s.OffCommand.Execute()
}

func main() {
	light := &BedroomLight{}

	switchOn := &TurnOnLightCommand{Light: light}
	switchOff := &TurnOffLightCommand{Light: light}

	lightSwitch := &Switch{
		OnCommand:  switchOn,
		OffCommand: switchOff,
	}

	lightSwitch.FlipOn()
	lightSwitch.FlipOff()
}

/*
В данном примере мы создали интерфейс `Light` для реализации конкретных получателей, в данном случае `BedroomLight`.
Затем мы создали интерфейс `Command`, который определяет метод `Execute()` для выполнения команды.
Реализовали два конкретных класса команд `TurnOnLightCommand` и `TurnOffLightCommand`, которые получают в качестве аргумента объект `Light`.
Наконец, мы создали класс `Switch`, который содержит две команды `OnCommand` и `OffCommand` и имеет методы `FlipOn()` и `FlipOff()`, которые вызывают метод `Execute()` соответствующей команды.

В главной функции мы создаем объект `BedroomLight`, объекты команд `TurnOnLightCommand` и `TurnOffLightCommand`, а также объект `Switch`, который использует эти команды.
Затем мы вызываем методы `FlipOn()` и `FlipOff()` для переключения света в комнате.
*/
