func minInsertions(s string) int {
    n := len(s)
    cnt := 0
    stack := make([]byte, 0)
    for i := 0; i < n; i++ {
        if s[i] == '(' {
            stack = append(stack, '(')
        } else {
            if len(stack) == 0 {
                stack = append(stack, '(')
                cnt++
            }
            ni := i+1
            if ni < n && s[ni] == ')' {
                i = ni
            } else {
                cnt++
            }
            stack = stack[0:len(stack)-1]
        }
    }
    return cnt + 2 * len(stack)
}
