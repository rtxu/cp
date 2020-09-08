func winnerSquareGame(n int) bool {
    state := make([]bool, n+1)
    state[0] = false
    availableMove := make([]int, 0)
    moveRoot := 1
    for i := 1; i <= n; i++ {
        for moveRoot * moveRoot <= i {
            availableMove = append(availableMove, moveRoot*moveRoot)
            moveRoot++
        }
        
        for j := 0; j < len(availableMove); j++ {
            if state[i - availableMove[j]] == false {
                // 对方必败
                state[i] = true
                break
            }
        }
    }
    
    return state[n]
}
