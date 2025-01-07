func tribonacci(n int) int {
    lastThreeSlice:= [3]int{0, 1, 1}
    if n < 3 {
        return lastThreeSlice[n]
    }

    for i := 3; i <= n; i++ {
        lastThreeSlice[i % 3] = lastThreeSlice[0] + lastThreeSlice[1] + lastThreeSlice[2]

    }

    return lastThreeSlice[n % 3]
}