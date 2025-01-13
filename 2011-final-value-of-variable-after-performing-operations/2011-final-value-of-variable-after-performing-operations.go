func finalValueAfterOperations(operations []string) int {
    x := 0

    for _, ops := range operations {
        if strings.HasPrefix(ops, "--") || strings.HasSuffix(ops, "--") {
            x--
        }

        if strings.HasPrefix(ops, "++") || strings.HasSuffix(ops, "++") {
            x++
        }
    }
    
    return x
}