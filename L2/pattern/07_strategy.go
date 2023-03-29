package main

import "fmt"

// Context
type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(a, b int) int {
	return c.strategy.Execute(a, b)
}

// Strategy interface
type Strategy interface {
	Execute(a, b int) int
}

// Concrete strategy
type ConcreteStrategyAdd struct{}

func (s *ConcreteStrategyAdd) Execute(a, b int) int {
	return a + b
}

// Concrete strategy
type ConcreteStrategySubtract struct{}

func (s *ConcreteStrategySubtract) Execute(a, b int) int {
	return a - b
}

func main() {
	context := &Context{}

	// Set the add strategy
	context.SetStrategy(&ConcreteStrategyAdd{})
	result := context.ExecuteStrategy(10, 5)
	fmt.Println(result)

	// Set the subtract strategy
	context.SetStrategy(&ConcreteStrategySubtract{})
	result = context.ExecuteStrategy(10, 5)
	fmt.Println(result)
}

/*
В данном примере мы создали класс `Context`, который использует стратегию для выполнения определенной операции.
Класс `Context` имеет методы `SetStrategy()` для установки стратегии и `ExecuteStrategy()` для выполнения операции с помощью текущей стратегии.

Затем мы создали интерфейс `Strategy`, который описывает метод `Execute()` для выполнения операции.
Мы также создали два конкретных класса стратегий `ConcreteStrategyAdd` и `ConcreteStrategySubtract`, которые реализуют метод `Execute()` интерфейса `Strategy`.
Каждый из этих классов выполняет свою операцию.

В главной функции мы создаем объект `Context` и устанавливаем в нем стратегию `ConcreteStrategyAdd`.
Затем мы вызываем метод `ExecuteStrategy()` объекта `Context`, передавая ему два аргумента, и выводим результат на экран.
Затем мы устанавливаем в объекте `Context` стратегию `ConcreteStrategySubtract` и снова вызываем метод `ExecuteStrategy()`, чтобы выполнить операцию вычитания и вывести результат на экран.
*/