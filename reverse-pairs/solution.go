func reversePairs(nums []int) int {
    return mergeSort(nums, 0, len(nums))
}

func mergeSort(nums []int, start, end int) int {
    if end - start <= 1 {
        return 0
    }
    mid := (start + end) >> 1
    count := mergeSort(nums, start, mid) + mergeSort(nums, mid, end)
    
    buf := make([]int, end-start)
    bufTail := 0
    j, k := mid, mid
    for i := start; i < mid; i++ {
        for k < end && 2*nums[k] < nums[i] {
            k++
        }
        for j < end && nums[j] < nums[i] {
            buf[bufTail] = nums[j]
            bufTail++
            j++
        }
        buf[bufTail] = nums[i]
        bufTail++
        
        count += k - mid
    }
    copy(nums[start:], buf[:bufTail])
    
    return count
}
