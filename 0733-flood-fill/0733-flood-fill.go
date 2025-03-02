func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	originalColor := image[sr][sc]
	if originalColor != color {
		dfs(image, sr, sc, originalColor, color)
	}
	return image
}

func dfs(image [][]int, r, c, originalColor, newColor int) {
	if image[r][c] == originalColor {
		image[r][c] = newColor
		if r >= 1 {
			dfs(image, r-1, c, originalColor, newColor)
		}
		if c >= 1 {
			dfs(image, r, c-1, originalColor, newColor)
		}
		if r+1 < len(image) {
			dfs(image, r+1, c, originalColor, newColor)
		}
		if c+1 < len(image[0]) {
			dfs(image, r, c+1, originalColor, newColor)
		}
	}
}