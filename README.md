# showfuzz
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
## Future Outlook
## License
The source code is licensed MIT. The website content is licensed CC BY 4.0,see LICENSE.
