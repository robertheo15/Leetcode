func maxWidthOfVerticalArea(points [][]int) int {
	tempArray := make([]int, 0)
	for i := 0; i < len(points); i++ {
		tempArray = append(tempArray, points[i][0])
	}

	maxWidth := 0
	sort.Ints(tempArray)

	for i := 0; i < len(tempArray)-1; i++ {
		if tempArray[i+1]-tempArray[i] > maxWidth {
			maxWidth = tempArray[i+1] - tempArray[i]
		}
	}

	return maxWidth
}