package edb

func main() {
	
}

// rotate2DSlice creates a 2D slice rotated 90 degrees clockwise
func rotate2DSlice(image [][]int) [][]int {
	newWidth := len(image)
	if newWidth == 0 {
		return image
	}
	newLength := len(image[0])

	newImage := make([][]int, newLength)
	for j := newLength - 1; j >= 0; j-- {
		row := make([]int, newWidth)
		for i := newWidth - 1; i >= 0; i-- {
			row[newWidth-1-i] = image[i][j]
		}
		newImage[j] = row
	}
	return newImage
}
