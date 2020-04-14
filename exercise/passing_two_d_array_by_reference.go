package exercise

//add2d  returns sum of 2d slice.
func add2d(arr *[][]int) (sum int) {
	rows := len(*(arr))
	for i := 0; i < rows; i++ {
		cols := len((*(arr))[i])
		for j := 0; j < cols; j++ {
			sum += (*arr)[i][j]
		}
	}

	return sum
}
