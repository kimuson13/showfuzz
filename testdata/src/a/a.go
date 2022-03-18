package a

type MyInt int

type Maybe struct {
	A int
	B error
}

type NG struct {
	F func(int, int)
}

func canFuzz1(a int) error { // want "can fuzz test"
	return nil
}

func canFuzz2(string, []byte) {} // want "can fuzz test"

func canFuzz3(float32, float64, bool) {} // want "can fuzz test"

func canFuzz4(MyInt) {} // want "can fuzz test"

func maybeFuzz(Maybe) {}

func cantFuzz1(error) {}

func cantFuzz2() {
	f := func(a int, b int) { // this is not func decl
		print(a, b)
	}

	f(1, 2)

	r := NG{f}
	r.F(1, 2)
}

func cantFuzz3() {}
