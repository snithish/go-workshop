package another_package

import "fmt"

const AnotherFizz = 0

const anotherFizz = 1

func userMainPackageGlobals() {
	fmt.Println(anotherFizz)
}
