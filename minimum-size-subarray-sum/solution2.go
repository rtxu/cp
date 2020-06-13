// Time: O(nlogn)
// Space: O(n)

func minSubArrayLen(s int, nums []int) int {
    n := len(nums)
    sum := make([]int, n+1)
    sum[0] = 0
    for i := 1; i <= n; i++ { sum[i] = sum[i-1]+ nums[i-1] }
    
    binarySearch := func(start, end int) int {
        if nums[start] >= s {
            return start+1
        }
        baseSum := 0
        if start > 0 {
            baseSum = sum[start]
        }
        if sum[end]-baseSum < s {
            return -1
        }
        
        for start+1 != end {
            mid := (start+end) >> 1
            if sum[mid]-baseSum >= s {
                end = mid
            } else {
                start = mid
            }
        }
        return end
    }

    bestLen := 0
    for i := 0; i < n; i++ {
        end := binarySearch(i, n)
        fmt.Println("i", i, "end", end)
        if end != -1 {
            if bestLen == 0 || end-i < bestLen {
               bestLen = end-i
            }
        }
    }
    return bestLen
}
