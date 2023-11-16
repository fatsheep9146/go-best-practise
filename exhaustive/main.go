package main

type Token int

const (
	Add Token = iota
	Subtract
	Multiply
	Quotient
	Remainder
)

func func1(t Token) {
	//exhaustive:ignore
	switch t {
	case Add:
	case Subtract:
	case Remainder:
	default:
	}
}

func func2(t Token) {
	//exhaustive:enforce
	switch t {
	case Add:
	case Subtract:
	case Remainder:
	default:
	}
}

func func3(t Token) {
	switch t {
	case Add:
	case Subtract:
	case Remainder:
	default:
	}
}

func main() {
	func1(Add)
	func2(Add)
}
