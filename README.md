# showfuzz [![Go Reference](https://pkg.go.dev/badge/github.com/kimuson13/showfuzz.svg)](https://pkg.go.dev/github.com/kimuson13/showfuzz)
showfuzz is report to stdout the functions that can be fuzzing test
## Caution
showfuzz is intended to be used for Go1.18 or higher
## Install
```
$ go install github.com/kimuson13/showfuzz/cmd/showfuzz@latest
```
## Situation to use
When you want to find a function that can fuzz test from existing Go codes,
`showfuzz` help that.
## How to use
If you want to check a package, give the package path of interest as the first argument:
```
$ showfuzz github.com/kimuson13/showfuzz
```
To check all packages beneath the current directory:
```
$ showfuzz ./...
```
## Demo
If you are on a directoty like that
```
$ tree .
sample
├── cmd
│   └── sample
│       └── main.go
├── go.mod
├── go.sum
└── sample.go
```
```
$ cat sample.go
package sample

func CanFuzzFunc(a int, b int) {
    return a * b
}
```
If you run `showfuzz`
```
$ showfuzz ./...
/GOPATH/sample/sample.go:3:6: CanFuzzFunc can fuzz test
```
## Future Outlook
### Don't report field values func
Now, `showfuzz` report field values func such as:
```
package main

type A struct {
    f func(int) int
}

func run (i int) { // run can fuzz test
    return i
}

func main() {
    a := A{
        f: run // run can fuzz test <- this report is not expected
    }
}
```
### Some Options
Now, `showfuzz` report all function that can fuzz test.  
If you already implement the fuzz test, you don't want to report that function.  
So, I want to add `ignore option`
## License
The source code is licensed MIT. The website content is licensed CC BY 4.0,see LICENSE.
