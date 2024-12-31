func numJewelsInStones(jewels string, stones string) int {
    freq := make(map[rune]int)
    
    for _, char := range jewels {
        freq[char]++
    }
    
    counter := 0

    for _, char := range stones {     
        if freq[char] == 1 {
            counter++
        }
    }
    
    return counter
}