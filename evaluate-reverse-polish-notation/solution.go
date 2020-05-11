func isNumber(token string) bool {
    tail := token[len(token) - 1]
    
    return '0' <= tail && tail <= '9'
}

func calc(left, right int, op byte) int {
    switch op {
    case '+':
        return left + right
    case '-':
        return left - right
    case '*':
        return left * right
    case '/':
        return left / right
    default:
        panic("unexpected operator: " + string(op))
    }
    
}

func evalRPN(tokens []string) int {
    argStack := make([]int, len(tokens))
    argCount := 0
    
    var arg int
    for i := 0; i < len(tokens); i++ {
        if isNumber(tokens[i]) {
            fmt.Sscanf(tokens[i], "%d", &arg)
            argStack[argCount] = arg
            argCount++
        } else {
            left, right := argStack[argCount-2], argStack[argCount-1]
            argCount -= 2
            argStack[argCount] = calc(left, right, tokens[i][0])
            argCount++
        }
    }
    return argStack[0]
}
