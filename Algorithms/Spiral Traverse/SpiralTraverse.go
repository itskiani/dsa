package main

func spiralTraverse(arr [][]int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	res := []int{}
	startRow, endRow := 0, len(arr)-1
	startCol, endCol := 0, len(arr[0])-1

	for startRow <= endRow && startCol <= endCol {

		for col := startCol; col <= endCol; col++ {
			res = append(res, arr[startRow][col])
		}

		for row := startRow + 1; row <= endRow; row++ {
			res = append(res, arr[row][endCol])
		}

		for col := endCol - 1; col >= startCol; col-- {
			if startRow == endRow {
				break
			}
			res = append(res, arr[endRow][col])
		}

		for row := endRow - 1; row > startRow; row-- {
			if startCol == endCol {
				break
			}
			res = append(res, arr[row][startCol])
		}

		startRow++
		endRow--
		startCol++
		endCol--
	}

	return res
}

func main() {

}
