package main

import (
	"context"
	"fmt"
	"github.com/d5/tengo/v2"
	"github.com/d5/tengo/v2/stdlib"
)

func main() {
	// create a new Script instance
	script := tengo.NewScript([]byte(
		`fmt := import("fmt")
each := func(seq, fn) {
    for x in seq { fn(x) }
}

sum := 0
mul := 1
each([a, b, c, d], func(x) {
	fmt.println("aa")
    sum += x
    mul *= x
})`))

	// set values
	_ = script.Add("a", 1)
	_ = script.Add("b", 9)
	_ = script.Add("c", 8)
	_ = script.Add("d", 4)

	script.SetImports(stdlib.GetModuleMap(stdlib.AllModuleNames()...))
	// run the script
	compiled, err := script.RunContext(context.Background())
	if err != nil {
		panic(err)
	}

	// retrieve values
	sum := compiled.Get("sum")
	mul := compiled.Get("mul")
	fmt.Println(sum, mul) // "22 288"
}
