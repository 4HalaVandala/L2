package main

import "fmt"

// Visitor is an interface for visiting different elements.
type Visitor interface {
	VisitElementA(elementA *ElementA)
	VisitElementB(elementB *ElementB)
}

// Element is an interface for accepting visitors.
type Element interface {
	Accept(visitor Visitor)
}

// ElementA is a concrete implementation of the Element interface.
type ElementA struct {
	name string
}

// Accept implements the Element interface for ElementA.
func (e *ElementA) Accept(visitor Visitor) {
	visitor.VisitElementA(e)
}

// ElementB is a concrete implementation of the Element interface.
type ElementB struct {
	value int
}

// Accept implements the Element interface for ElementB.
func (e *ElementB) Accept(visitor Visitor) {
	visitor.VisitElementB(e)
}

// ConcreteVisitor is a concrete implementation of the Visitor interface.
type ConcreteVisitor struct{}

// VisitElementA implements the Visitor interface for ConcreteVisitor.
func (v *ConcreteVisitor) VisitElementA(elementA *ElementA) {
	fmt.Printf("Visited ElementA with name %s\n", elementA.name)
}

// VisitElementB implements the Visitor interface for ConcreteVisitor.
func (v *ConcreteVisitor) VisitElementB(elementB *ElementB) {
	fmt.Printf("Visited ElementB with value %d\n", elementB.value)
}

func main() {
	elements := []Element{
		&ElementA{name: "element A"},
		&ElementB{value: 42},
	}

	visitor := &ConcreteVisitor{}

	for _, element := range elements {
		element.Accept(visitor)
	}
}

/*
В этом примере мы создаем интерфейс Посетителя (Visitor), который определяет методы для посещения различных элементов.
Затем мы создаем интерфейс Элемента (Element), который определяет метод для принятия посетителя.

Мы создаем конкретные реализации Элемента (ElementA и ElementB), которые реализуют метод Accept для принятия посетителя.
Мы также создаем конкретную реализацию Посетителя (ConcreteVisitor), которая реализует методы VisitElementA и VisitElementB для посещения соответствующих элементов.


Клиентский код может добавлять новые элементы, не затрагивая код Посетителя, и наоборот - добавлять новые методы Посетителя, не затрагивая код Элементов.
*/
