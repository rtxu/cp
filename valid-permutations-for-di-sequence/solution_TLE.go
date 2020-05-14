
const MOD = 1e9 + 7

func _numPermsDISequence(S string, M map[string]int) int {

    if _, exist := M[S]; exist {
        return M[S]
    }
    
       
    if len(S) == 1 {
        M[S] = 1
        return M[S]
    }
    
    count := 0
    n := len(S) + 1
    for i := 0; i < n; i++ {
        // 将 n 放在排列的第 i 个位置
        if i == 0 {
            if S[0] == 'D' {
                count += _numPermsDISequence(S[1:], M)
                count %= MOD
            }
        } else if i == n-1 {
            if S[n-2] == 'I' {
                count += _numPermsDISequence(S[0:n-2], M)
                count %= MOD
            }
        } else {
            if S[i-1] == 'I' && S[i] == 'D' {
                count += _numPermsDISequence(S[0:i-1] + "I" + S[i+1:], M)
                count += _numPermsDISequence(S[0:i-1] + "D" + S[i+1:], M)
                count %= MOD
            }
        }
    }
    
    M[S] = count
    return M[S]
}

func numPermsDISequence(S string) int {
    M := make(map[string]int)
    
    return _numPermsDISequence(S, M)
}
