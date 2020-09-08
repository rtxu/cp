func countRangeSum(nums []int, lower int, upper int) int {
    n := len(nums)
    sum := make([]int, n+1)
    for i := 0; i < n; i++ {
        sum[i+1] = sum[i] + nums[i]
    }
    return mergeSort(sum, 0, n+1, lower, upper)
}

func mergeSort(sum []int, start, end int, lower, upper int) int {
    if end-start <= 1 {
        return 0
    }
    mid := start + (end-start) >> 1
    count := mergeSort(sum, start, mid, lower, upper) + mergeSort(sum, mid, end, lower, upper)
    
    buf, bufTail := make([]int, end-start), 0
    j, k, r := mid, mid, mid
    for i := start; i < mid; i++ {
        for j < end && sum[j] - sum[i] < lower { j++ }
        for k < end && sum[k] - sum[i] <= upper { k++ }
        count += k - j
        
        for r < end && sum[r] < sum[i] {
            buf[bufTail] = sum[r]
            bufTail++
            r++
        }
        buf[bufTail] = sum[i]
        bufTail++
    }
    copy(sum[start:], buf[:bufTail])
    return count
}
