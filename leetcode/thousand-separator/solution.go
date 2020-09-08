func thousandSeparator(n int) string {
    if n == 0 {
        return "0"
    }
    stack := make([]int, 0)
    for n > 0 {
        stack = append(stack, n % 1000)
        n /= 1000
    }
    var result strings.Builder
    for i := len(stack)-1; i >= 0; i-- {
        if i < len(stack)-1 {
            result.WriteString(fmt.Sprintf("%03d", stack[i]))
        } else {
            result.WriteString(fmt.Sprintf("%d", stack[i]))
        }
        if i > 0 {
            result.WriteByte('.')
        }
    }
    return result.String()
}
