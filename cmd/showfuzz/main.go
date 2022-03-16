package main

import (
	"github.com/kimuson13/showfuzz"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(showfuzz.Analyzer) }
