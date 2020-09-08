// Time: O(n)
// Space: O(1)

func longestOnes(A []int, K int) int {
    var count [2]int
    
    maxLen := 0
    for begin, end := 0, 0; end < len(A); end++ {
        count[A[end]]++
        
        for count[0] > K {
            count[A[begin]]--
            begin++
        }
        if end-begin+1 > maxLen {
            maxLen = end-begin+1
        }
    }
    return maxLen
}
