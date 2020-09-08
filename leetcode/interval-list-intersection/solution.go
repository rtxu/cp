func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func intervalIntersection(A [][]int, B [][]int) [][]int {
    nA, nB := len(A), len(B)
    iA, iB := 0, 0
    var result [][]int
    for iA < nA && iB < nB {
        var left, right []int
        if A[iA][0] < B[iB][0] || (A[iA][0] == B[iB][0] && A[iA][1] < B[iB][1]) {
            left, right = A[iA], B[iB]
        } else {
            left, right = B[iB], A[iA]
        }
        if left[1] >= right[0] {
            result = append(result, []int{right[0], min(right[1], left[1])})
        }
        
        if A[iA][1] < B[iB][1] {
            iA++
        } else if A[iA][1] > B[iB][1] {
            iB++
        } else {
            iA++
            iB++
        }
    }
    
    return result
}
