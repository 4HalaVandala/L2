package main

import "fmt"

// Context
type Contexts struct {
	state State
}

func (c *Contexts) SetState(state State) {
	c.state = state
}

func (c *Contexts) Request() {
	c.state.Handle()
}

// State interface
type State interface {
	Handle()
}

// Concrete state
type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle() {
	fmt.Println("Handling state A")
}

// Concrete state
type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle() {
	fmt.Println("Handling state B")
}

func main() {
	context := &Contexts{}

	// Set the initial state to A
	context.SetState(&ConcreteStateA{})
	context.Request()

	// Change the state to B
	context.SetState(&ConcreteStateB{})
	context.Request()
}

/*
В данном примере мы создали класс `Context`, который содержит текущее состояние объекта.
Класс `Context` имеет методы `SetState()` для установки состояния и `Request()` для выполнения операции в зависимости от текущего состояния.

Затем мы создали интерфейс `State`, который описывает метод `Handle()` для выполнения операции в зависимости от текущего состояния.
Мы также создали два конкретных класса состояний `ConcreteStateA` и `ConcreteStateB`, которые реализуют метод `Handle()` интерфейса `State`.
Каждый из этих классов выполняет свою операцию в зависимости от текущего состояния.

В главной функции мы создаем объект `Context` и устанавливаем в нем начальное состояние `ConcreteStateA`.
Затем мы вызываем метод `Request()` объекта `Context`, который выполняет операцию в зависимости от текущего состояния.
Затем мы устанавливаем в объекте `Context` состояние `ConcreteStateB` и снова вызываем метод `Request()`, чтобы выполнить операцию в зависимости от нового состояния.
*/
