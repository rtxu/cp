func sign(diff int) int {
    if diff == 0 {
        return 0
    } else if diff > 0 {
        return 1
    } else {
        return -1
    }
}

func updateIfGreater(cur int, max *int) {
    if cur > *max {
        *max = cur
    }
}

func maxTurbulenceSize(A []int) int {
    if len(A) < 2 {
        return len(A)
    }
    
    var maxSeqLen, curSeqLen, lastSign int
    for i := 1; i < len(A); i++ {
        currentSign := sign(A[i] - A[i-1])
        seqSign := lastSign * currentSign
        
        switch seqSign {
        case 0:
            if currentSign == 0 {
                curSeqLen = 1
            } else {
                curSeqLen = 2
            }
        case 1:
            curSeqLen = 2
        case -1:
            curSeqLen++
        }
        updateIfGreater(curSeqLen, &maxSeqLen)
        lastSign = currentSign
    }
    
    return maxSeqLen
}
