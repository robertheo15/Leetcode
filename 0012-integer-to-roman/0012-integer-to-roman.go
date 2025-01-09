func intToRoman(num int) string {
	roman := []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	romanString := ""

	for _, entry := range roman {
		for num >= entry.value {
			romanString += entry.symbol
			num -= entry.value
		}
	}

	return romanString
}