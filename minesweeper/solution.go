func revealed(b byte) bool {
    return b != 'E' && b != 'M'
}

var dx = []int{-1, -1, -1, 0, 1, 1, 1, 0}
var dy = []int{-1, 0, 1, 1, 1, 0, -1, -1}

func dfs(board [][]byte, x, y int) {
    n, m := len(board), len(board[0])
    if x < 0 || x >= n || y < 0 || y >= m {
        return
    }
    if revealed(board[x][y]) {
        return
    }
    
    mineCount := 0
    for i := 0; i < 8; i++ {
        nx, ny := x + dx[i], y + dy[i]
        if nx >= 0 && nx < n && ny >= 0 && ny < m {
            if board[nx][ny] == 'M' {
                mineCount++
            }
        }
    }
    if mineCount > 0 {
        board[x][y] = '0' + byte(mineCount)
    } else {
        board[x][y] = 'B'
        for i := 0; i < 8; i++ {
            nx, ny := x + dx[i], y + dy[i]
            dfs(board, nx, ny)
        }
    }
}

func updateBoard(board [][]byte, click []int) [][]byte {
    if board[click[0]][click[1]] == 'M' {
        board[click[0]][click[1]] = 'X'
        return board
    } else {
        dfs(board, click[0], click[1])
        return board
    }
}
