func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	if haystack == needle {
		return 0
	}
    
	h := len(haystack)
	n := len(needle)

	if n > h {
		return -1
	}

	for i := 0; i < h-n+1; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}

	return -1
}
