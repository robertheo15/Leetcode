func fizzBuzz(n int) []string {
	output := make([]string, 0)

	for i := 1; i <= n; i++ {
		answer := ""

		if i % 15 == 0 { 
			answer = "FizzBuzz"
		} else if i % 3 == 0  { 
			answer = "Fizz"
		} else if i % 5 == 0  { 
			answer = "Buzz"
		} else {
			answer = fmt.Sprintf("%d", i)
		}

		output = append(output, answer)
	}

	return output
}