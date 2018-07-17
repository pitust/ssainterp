package ssainterp_test

import (
	"bytes"
	"fmt"

	"github.com/pitust/ssainterp"
)

const code = `
package main

import "fmt"

func main() {
	fmt.Println("42")
}

`

// ExampleRun trivial example, more to add.
func ExampleRun() {
	var output bytes.Buffer
	ssainterp.Run(code, nil, nil, &output)
	fmt.Println(output.String())
	// Output: 42
}
