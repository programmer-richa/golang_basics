package exercise

import "github.com/programmer-richa/golang_basics/utility"

// Swap swaps 2 numbers simply by x,y=y,X Operation.
func Swap() {
	x, y := 10, 20
	utility.Println(sublevel, "X: ", x, " y: ", y)
	x, y = y, x
	utility.Println(sublevel, "After swapping")
	utility.Println(sublevel, "X: ", x, " y: ", y)
}
