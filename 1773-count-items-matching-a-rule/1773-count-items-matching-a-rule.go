func countMatches(items [][]string, ruleKey string, ruleValue string) int {
    result := 0
    sliceKey := []string{"type", "color", "name"}

    for _, item := range items {
        for i, key := range sliceKey{
            if ruleKey == key && item[i] == ruleValue {
                result++
            }
        }
    }

    return result
}

