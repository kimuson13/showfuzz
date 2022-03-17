package a

type MyInt int

type Maybe struct {
	A int
	B error
}

func canFuzz1(a int) error { // want "can fuzz test"
	return nil
}

func canFuzz2(string, []byte) {} // want "can fuzz test"

func canFuzz3(float32, float64, bool) {} // want "can fuzz test"

func canFuzz4(MyInt) {} // want "can fuzz test"

func maybeFuzz(Maybe) {}

func cantFuzz1(error) {}

func cantFuzz3() {}
