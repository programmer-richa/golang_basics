package exercise

import "github.com/programmer-richa/golang_basics/utility"

//swap   swap 2 numbers simply by x,y=y,X Operation.
func swap() {
	x, y := 10, 20
	utility.Println(sublevel, "X: ", x, " y: ", y)
	x, y = y, x
	utility.Println(sublevel, "After swapping")
	utility.Println(sublevel, "X: ", x, " y: ", y)
}
