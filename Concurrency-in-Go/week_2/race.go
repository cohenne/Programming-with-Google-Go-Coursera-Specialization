package main

import (
	"fmt"
)

func main() {
	for {
		a := 0
		b := 0

		go func(x *int) {
			*x = 1
		}(&a)

		go func(x *int) {
			*x = 1
		}(&b)

		go func(x int, y int) {
			fmt.Println(x / y)
		}(a, b)

	}
}

/*
A race condition occurs when multiple threads (goroutines here)
are reading and writing to the same variables.

In the example below, we are trying to divide a by b.
However, if the divison goroutine occurs before a and b are set,
a divide-by-zero exception will occur.
Because the goroutine scheduling mechanism is non-deterministic,
 we are likely to run into this exception as we iterate through the for loop.
*/
