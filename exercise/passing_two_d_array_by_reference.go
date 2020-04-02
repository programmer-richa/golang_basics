package exercise

//AddElementsOf2DArray : Returns sum of 2d array
func AddElementsOf2DArray(arr *[][]int) int {
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
