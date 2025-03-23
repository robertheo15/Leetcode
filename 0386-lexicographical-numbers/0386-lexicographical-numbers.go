func lexicalOrder(n int) []int {
    dictionary := make([]string, 0)

    for i := 1; i <= n ; i++ {
        dictionary = append(dictionary, strconv.Itoa(i))
    }

    sort.Strings(dictionary)
    result := make([]int, 0)

    for _, val := range dictionary {
        num, _ := strconv.Atoi(val)
        result = append(result, num)
    }

    return result
}