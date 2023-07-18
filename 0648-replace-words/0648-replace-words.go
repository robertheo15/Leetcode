func replaceWords(dictionary []string, sentence string) string {
    brokenString := strings.Fields(sentence)
    results := make([]string, 0)

    for _, val := range brokenString {
        var tempRes string
        for _, dicVal := range dictionary {
            if strings.HasPrefix(val, dicVal) {
                if len(tempRes) == 0 {
                    tempRes = dicVal
                } else if len(dicVal) < len(tempRes) {
                    tempRes = dicVal
                }
            }
        }

        if len(tempRes) > 0 {
             results = append(results, tempRes)
        }else{
            results = append(results, val)
        }
    }

    return strings.Join(results, " ")

}