package a

func f() {
	// The pattern can be written in regular expression.
	var gopher int // want "pattern"
	print(gopher)  // want "identifier is gopher"
}

func canFuzz1(int) error {} // want "can fuzz test"

func canFuzz2(string, []byte) {} // want "can fuzz test"

func canFuzz3(float32, float64, bool) {} // want "can fuzz test"

func cantFuzz1(error) {}

func cantFuzz2(any) {}
