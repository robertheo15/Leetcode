func sortPeople(names []string, heights []int) []string {
	mapNames := make(map[int]string)
	for i := 0; i < len(names); i++ {
		mapNames[heights[i]] = names[i]
	}

	sort.Ints(heights)

	for i := range heights {
		names[len(names)-1-i] = mapNames[heights[i]]
	}

	return names
}