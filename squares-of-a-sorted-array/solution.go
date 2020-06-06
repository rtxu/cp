func sortedSquares(A []int) []int {
    n := len(A)
    var ge0 int // the first >= 0 index
    for ge0 = 0; ge0 < n && A[ge0] < 0; ge0++ {}
    
    result := make([]int, n)
    pos, neg := ge0, ge0-1
    for i := 0; i < n; i++ {
        posV, negV := -1, -1
        if pos < n {
            posV = A[pos] * A[pos]
        }
        if neg >= 0 {
            negV = A[neg] * A[neg]
        }
        fmt.Println(i, pos, posV, neg, negV)
        if posV == -1 {
            result[i] = negV
            neg--
        } else if negV == -1 {
            result[i] = posV
            pos++
        } else {
            if posV < negV {
                result[i] = posV
                pos++
            } else {
                result[i] = negV
                neg--
            }
        }
    }
    
    return result
}
