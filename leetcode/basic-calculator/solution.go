

func isNumber(b byte) bool {
    return '0' <= b && b <= '9'
}

func calc(left, right int, op byte) int {
    switch op {
    case '+':
        return left + right
    case '-':
        return left - right
    }
    panic("unexpected op: " + string(op))
}

func calculate(s string) int {
    opStack := make([]byte, len(s))
    argStack := make([]interface{}, len(s))
    opCount, argCount := 0, 0
    
    for i := 0; i < len(s); i++ {
        if isNumber(s[i]) {
            num := int(s[i] - '0')
            for i++; i < len(s) && isNumber(s[i]); i++ {
                num *= 10
                num += int(s[i] - '0')
            }
            i--
            argStack[argCount] = num
            argCount++
        } else if s[i] == '+' || s[i] == '-' {
            opStack[opCount] = s[i]
            opCount++
        } else if s[i] == '(' {
            argStack[argCount] = s[i]
            argCount++
        } else if s[i] == ')' {
            // argCount - 2 is LEFT_PARENTHESE
            argStack[argCount - 2] = argStack[argCount - 1]
            argCount--
        } else {
            // ignore
        }
        
        for opCount > 0 && argCount >= 2 {
            left, leftOk := argStack[argCount-2].(int)
            right, rightOk := argStack[argCount-1].(int)
            if leftOk && rightOk {
                argCount -= 2
                argStack[argCount] = calc(left, right, opStack[opCount-1])
                opCount--
                argCount++
            } else {
                break
            }
        }
    }
    
    return argStack[0].(int)
}
