package main

func foo(f Fooer) {
	f.Foo()
}

func main() {
	f := NewFooer()
	foo(f)
}
