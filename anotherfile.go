package main

import "fmt"

type Fooer interface {
	Foo()
}

type fooer struct{}

func NewFooer() *fooer {
	return &fooer{}
}

func (f *fooer) Foo() {
	fmt.Println("foo")
}
