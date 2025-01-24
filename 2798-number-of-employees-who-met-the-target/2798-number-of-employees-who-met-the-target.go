func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	counter := 0

	for _, hour := range hours {
		if hour >= target {
			counter++
		}
	}

	return counter
}