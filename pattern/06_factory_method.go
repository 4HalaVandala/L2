package main

import "fmt"

// Product interface
type Products interface {
	GetName() string
}

// Concrete product
type ConcreteProductA struct{}

func (p *ConcreteProductA) GetName() string {
	return "ConcreteProductA"
}

// Concrete product
type ConcreteProductB struct{}

func (p *ConcreteProductB) GetName() string {
	return "ConcreteProductB"
}

// Creator interface
type Creator interface {
	Create() Products
}

// Concrete creator
type ConcreteCreatorA struct{}

func (c *ConcreteCreatorA) Create() Products {
	return &ConcreteProductA{}
}

// Concrete creator
type ConcreteCreatorB struct{}

func (c *ConcreteCreatorB) Create() Products {
	return &ConcreteProductB{}
}

func main() {
	creatorA := &ConcreteCreatorA{}
	productA := creatorA.Create()
	fmt.Println(productA.GetName())

	creatorB := &ConcreteCreatorB{}
	productB := creatorB.Create()
	fmt.Println(productB.GetName())
}

/*
В данном примере мы создали интерфейс `Product` для описания продукта, который будет создаваться фабричным методом.
Затем мы создали два конкретных класса продуктов `ConcreteProductA` и `ConcreteProductB`, которые реализуют метод `GetName()` интерфейса `Product`.

Далее мы создали интерфейс `Creator` для описания фабричного метода, который будет создавать продукты.
Затем мы создали два конкретных класса фабричных методов `ConcreteCreatorA` и `ConcreteCreatorB`, которые реализуют метод `Create()` интерфейса `Creator`.
Каждый фабричный метод создает свой конкретный продукт.

В главной функции мы создаем объекты каждого из двух фабричных методов и вызываем их метод `Create()`, чтобы создать соответствующие продукты.
Затем мы выводим на экран имена созданных продуктов.
*/
