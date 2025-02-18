func countMatches(items [][]string, ruleKey string, ruleValue string) int {
    result := 0

    for _, item := range items {
        if ruleKey == "type" && item[0] == ruleValue {
			result++
		}
		if ruleKey == "color" && item[1] == ruleValue {
			result++
		} else if ruleKey == "name" && item[2] == ruleValue {
			result++
		}
        // for _, i := range item {
        //     if ruleKey == "type" && i == ruleValue {
        //         result +=1
        //     } else if ruleKey == "color" && i== ruleValue{
        //         result +=1
        //     } else if ruleKey == "name" && i == ruleValue{
        //         result +=1
        //     }
        // }
    }

    return result
}

