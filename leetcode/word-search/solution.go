// Time: O(4*row*col)
func backtrack(x, y, matchedLen int, visited [][]bool, board [][]byte, word string) bool {
    if matchedLen == len(word) {
        return true
    }
    if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
        return false
    }
    if visited[x][y] {
        return false
    }

    if board[x][y] == word[matchedLen] {
        matchedLen++
        visited[x][y] = true
        if backtrack(x-1, y, matchedLen, visited, board, word) {
            return true
        }
        if backtrack(x+1, y, matchedLen, visited, board, word) {
            return true
        }
        if backtrack(x, y-1, matchedLen, visited, board, word) {
            return true
        }
        if backtrack(x, y+1, matchedLen, visited, board, word) {
            return true
        }
        matchedLen--
        visited[x][y] = false
        return false
    } else {
        return false
    }
}

func exist(board [][]byte, word string) bool {
    row := len(board)
    if row == 0 {
        return len(word) == 0
    }
    col := len(board[0])
    visited := make([][]bool, row)
    for i := 0; i < row; i++ {
        visited[i] = make([]bool, col)
    }
    for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            if backtrack(i, j, 0, visited, board, word) {
                return true
            }
        }
    }
    return false
}
