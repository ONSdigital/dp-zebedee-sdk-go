// This is intentionally badly formatted code to test the linter in CI.
package main

import "fmt"

var (
	ThisIsAPublicVarWithoutGodoc string

	errSomeError error
)

func main() {
	var s string = "Hello world!"

	fmt.Println(s)
}

func DoSomething() int {
	if true {
		return 0
		return 0  // intentionally adding this to trigger vet failure.
	}

	return 1
}
