package main

import "fmt"

// Handler interface
type Handler interface {
	SetNext(handler Handler) Handler
	Handle(request string) string
}

// Concrete handler
type LowercaseHandler struct {
	next Handler
}

func (h *LowercaseHandler) SetNext(handler Handler) Handler {
	h.next = handler
	return handler
}

func (h *LowercaseHandler) Handle(request string) string {
	if request == "lowercase" {
		return "Handled by LowercaseHandler"
	} else if h.next != nil {
		return h.next.Handle(request)
	} else {
		return "Request not handled"
	}
}

// Concrete handler
type UppercaseHandler struct {
	next Handler
}

func (h *UppercaseHandler) SetNext(handler Handler) Handler {
	h.next = handler
	return handler
}

func (h *UppercaseHandler) Handle(request string) string {
	if request == "uppercase" {
		return "Handled by UppercaseHandler"
	} else if h.next != nil {
		return h.next.Handle(request)
	} else {
		return "Request not handled"
	}
}

// Concrete handler
type NumberHandler struct {
	next Handler
}

func (h *NumberHandler) SetNext(handler Handler) Handler {
	h.next = handler
	return handler
}

func (h *NumberHandler) Handle(request string) string {
	if request == "number" {
		return "Handled by NumberHandler"
	} else if h.next != nil {
		return h.next.Handle(request)
	} else {
		return "Request not handled"
	}
}

func main() {
	lowercase := &LowercaseHandler{}
	uppercase := &UppercaseHandler{}
	number := &NumberHandler{}

	lowercase.SetNext(uppercase).SetNext(number)

	fmt.Println(lowercase.Handle("lowercase"))
	fmt.Println(lowercase.Handle("uppercase"))
	fmt.Println(lowercase.Handle("number"))
	fmt.Println(lowercase.Handle("unknown"))
}

/*
В данном примере мы создали интерфейс `Handler` для реализации конкретных обработчиков запросов.
Затем мы создали три конкретных класса обработчиков `LowercaseHandler`, `UppercaseHandler` и `NumberHandler`, которые реализуют методы интерфейса `Handler`.
Каждый обработчик может либо обработать запрос, либо передать его следующему обработчику в цепочке.

В главной функции мы создаем объекты каждого из трех обработчиков и устанавливаем связь между ними с помощью метода `SetNext()`.
Затем мы вызываем метод `Handle()` у первого обработчика в цепочке, передавая ему различные запросы.
Каждый обработчик либо обрабатывает запрос, либо передает его следующему обработчику в цепочке.
Если запрос не может быть обработан ни одним из обработчиков, будет возвращена строка "Request not handled".
*/