package main

import "fmt"

// Subsystem1 is a part of the complex system.
type Subsystem1 struct{}

// Operation1 is a method of Subsystem1.
func (s *Subsystem1) Operation1() {
	fmt.Println("Subsystem1 operation 1")
}

// Operation2 is another method of Subsystem1.
func (s *Subsystem1) Operation2() {
	fmt.Println("Subsystem1 operation 2")
}

// Subsystem2 is another part of the complex system.
type Subsystem2 struct{}

// Operation3 is a method of Subsystem2.
func (s *Subsystem2) Operation3() {
	fmt.Println("Subsystem2 operation 3")
}

// Operation4 is another method of Subsystem2.
func (s *Subsystem2) Operation4() {
	fmt.Println("Subsystem2 operation 4")
}

// Facade provides a simplified interface to the complex system.
type Facade struct {
	subsystem1 *Subsystem1
	subsystem2 *Subsystem2
}

// NewFacade creates a new Facade object.
func NewFacade() *Facade {
	return &Facade{
		subsystem1: &Subsystem1{},
		subsystem2: &Subsystem2{},
	}
}

// OperationWrapper provides a simplified interface to the complex system.
func (f *Facade) OperationWrapper() {
	f.subsystem1.Operation1()
	f.subsystem2.Operation3()
}

func main() {
	facade := NewFacade()
	facade.OperationWrapper()
}

/*
Фасад скрывает детали реализации подсистем и предоставляет только необходимый функционал.
В нашем примере фасад предоставляет только один метод (OperationWrapper), который вызывает методы подсистем (Operation1 и Operation3).

*/
