package main

import (
	"fmt"
	"go/types"
	"programming_constructs/another_package"
)

const BUZZ string = "Buzz"

const FIZZ = "Fizz" // Types are optional owing to type inference

var fizzGlobal, buzzGlobal = "fizz", "buzz" // multi variable declaration is possible

var (
	fizzMultiVar = "fizz" // multi line var declarations in global scope
	buzzMultiVar = "buzz"
)

func main() {
	unusedVar := "unused" // compilation will fail if there are unused variables

	// fizzLocal = "FIZZ" -> fails := to be used for declaration inside blocks

	fizzLocal := "FIZZ"     // declaration plus initialization
	fizzLocal = "localFizz" // mutating does not require := mutation does not mean usage

	useGlobals()
	if true {
		fmt.Println(another_package.AnotherFizz)
		fmt.Println(another_package.anotherFizz) // unexported like private scope
	}
}
