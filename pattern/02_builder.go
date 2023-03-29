package main

import "fmt"

// Product is the final object that the builder creates.
type Product struct {
	part1 string
	part2 int
}

// Builder is an interface for creating the product.
type Builder interface {
	BuildPart1()
	BuildPart2()
	GetProduct() *Product
}

// ConcreteBuilder is a concrete implementation of the Builder interface.
type ConcreteBuilder struct {
	product *Product
}

// NewConcreteBuilder creates a new ConcreteBuilder object.
func NewConcreteBuilder() *ConcreteBuilder {
	return &ConcreteBuilder{
		product: &Product{},
	}
}

// BuildPart1 implements the Builder interface for ConcreteBuilder.
func (b *ConcreteBuilder) BuildPart1() {
	b.product.part1 = "part 1"
}

// BuildPart2 implements the Builder interface for ConcreteBuilder.
func (b *ConcreteBuilder) BuildPart2() {
	b.product.part2 = 42
}

// GetProduct implements the Builder interface for ConcreteBuilder.
func (b *ConcreteBuilder) GetProduct() *Product {
	return b.product
}

// Director is responsible for using the builder to create the product.
type Director struct {
	builder Builder
}

// NewDirector creates a new Director object.
func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

// Construct builds the product using the builder.
func (d *Director) Construct() *Product {
	d.builder.BuildPart1()
	d.builder.BuildPart2()
	return d.builder.GetProduct()
}

func main() {
	builder := NewConcreteBuilder()
	director := NewDirector(builder)
	product := director.Construct()

	fmt.Printf("%+v\n", product)
}

/*
Мы создаем продукт (Product), который состоит из двух частей (part1 и part2).
Затем мы создаем интерфейс Строителя (Builder), который определяет методы для создания продукта.

Мы создаем конкретную реализацию Строителя (ConcreteBuilder), которая реализует методы интерфейса Builder.
Мы также создаем Директора (Director), который использует Строителя для создания продукта.

Мы видим, что мы успешно создали продукт с помощью паттерна Строитель.
Клиентский код не знает о деталях реализации Строителя и Директора, и может использовать их для создания продукта без знания о том, как он создается.
*/
