func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    result := false

    s1 := strings.Join(word1, "")
    s2 := strings.Join(word2, "")

    if s1 == s2 {
        return true
    }

    return result
}