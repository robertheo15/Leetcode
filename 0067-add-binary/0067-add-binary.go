func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	indexB := len(b) - 1
	result := make([]byte, len(a))

	var temp, sum byte
	for i := len(a) - 1; i >= 0; i-- { // backward iterate
		sum = temp + a[i] // sum a[i]
		if indexB >= 0 {
			sum += b[indexB] // sum b[i]
			indexB--
		}

		result[i] = sum%2 + '0'
		temp = sum >> 1 % 2 // (sun / 2, 1 time) % 2
	}
    
	if temp == 0 {
		return string(result)
	}

	return "1" + string(result)
}