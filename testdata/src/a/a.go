package a

type MyInt int

func canFuzz1(int) error { // want "can fuzz test"
	return nil
}

func canFuzz2(string, []byte) {} // want "can fuzz test"

func canFuzz3(float32, float64, bool) {} // want "can fuzz test"

func canFuzz4(MyInt) {} // want "can fuzz test"

func cantFuzz1(error) {}

func cantFuzz2(any) {}

func cantFuzz3() {}
