func next(i, j, n int) (nextI int, nextJ int) {
    nextI, nextJ = j, n-1-i
    return nextI, nextJ
}

func rotate(matrix [][]int)  {
    n := len(matrix)
    for i := 0; i < n; i++ {
        for j := i; j < n-1-i; j++ {
            currentI, currentJ := i, j
            currentV := matrix[currentI][currentJ]
            for k := 0; k < 4; k++ {
                nextI, nextJ := next(currentI, currentJ, n)
                nextV := matrix[nextI][nextJ]
                matrix[nextI][nextJ] = currentV
                currentI, currentJ = nextI, nextJ
                currentV = nextV
            }
        }
    }
}
