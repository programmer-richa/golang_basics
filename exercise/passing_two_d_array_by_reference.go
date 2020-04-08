package exercise

//add2d : Returns sum of 2d array
func add2d(arr *[][]int) int {
	var sum int = 0
	rows := len(*(arr))
	for i := 0; i < rows; i++ {
		cols := len((*(arr))[i])
		for j := 0; j < cols; j++ {
			sum += (*arr)[i][j]
		}
	}

	return sum
}
