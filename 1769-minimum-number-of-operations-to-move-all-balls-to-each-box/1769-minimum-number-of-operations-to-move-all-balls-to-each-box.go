func minOperations(boxes string) []int {
    indices := []int{}
    for i := 0; i < len(boxes); i++ {
        if boxes[i] == '1' {
            indices = append(indices, i)
        }
    }

    result := make([]int, len(boxes))
    for i := 0; i < len(boxes); i++ {
        cnt := 0
        for _, index := range indices {
            cnt += abs(index - i)
        }
        result[i] = cnt
    }
    return result
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}