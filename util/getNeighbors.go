package util

// Gets all possible neighbors of a position within a 2D slice
//
// arr is 2d array of any type, x and y are integers representing position of center element
// elements are returned in order:
//
// [1,2,3]
//
// [3,x,4]
//
// [5,6,7]
//
// x will be a zero value, unless a specific direction is passed, then it will be an element neighboring center from the selected side
func GetNeighborsDiagonal[T any](arr [][]T, x, y int, selector *[2]int) [3][3]T {

	neighbors := [3][3]T{}

	//you are outside of the array, you don't exits, so you don't have neighbors
	if x > len(arr[0])-1 || y > len(arr)-1 || x < 0 || y < 0 {
		return neighbors
	}

	if y > 0 {
		if x > 0 {
			neighbors[0][0] = arr[y-1][x-1]
		}
		neighbors[0][1] = arr[y-1][x]
		if x < len(arr[0])-1 {
			neighbors[0][2] = arr[y-1][x+1]
		}
	}

	if x > 0 {
		neighbors[1][0] = arr[y][x-1]
	}

	if x < len(arr[0])-1 {
		neighbors[1][2] = arr[y][x+1]
	}

	if y < len(arr)-1 {
		if x > 0 {
			neighbors[2][0] = arr[y+1][x-1]
		}
		neighbors[2][1] = arr[y+1][x]
		if x < len(arr[0])-1 {
			neighbors[2][2] = arr[y+1][x+1]
		}
	}

	if selector != nil {
		sx := selector[0]
		sy := selector[1]

		if 1+sy >= 0 && 1+sx >= 0 && 1+sx < len(arr)-1 && 1+sy < len(arr)-1 {
			neighbors[1][1] = neighbors[1+sy][1+sx]
		}

	}

	return neighbors
}
