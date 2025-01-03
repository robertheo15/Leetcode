func construct2DArray(original []int, m int, n int) [][]int {
    if len(original) != m*n {
        return [][]int{}
    }

    results := make([][]int, m)
    for i := 0; i < m; i++ {
        results[i] = make([]int, n)
        for j := 0; j < n; j++ {
            results[i][j] = original[i*n+j]
        }
    }

    return results
}