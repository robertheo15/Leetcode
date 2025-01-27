func minMovesToSeat(seats []int, students []int) int {
	counter := 0
	sort.Ints(seats)
	sort.Ints(students)

	for i, seat := range seats {
		result := abs(seat - students[i])

		counter += result
	}

	return counter
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}