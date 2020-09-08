func upper_bound(nums []int, target int) int {
    n := len(nums)
    if n == 0 || nums[n-1] <= target {
        return n
    }
    
    l, r := 0, n-1
    for l < r {
        mid := l + (r-l)>>1
        if nums[mid] > target {
            r = mid
        } else {
            // Punchline: due to mid is biased to left, left = mid will make for-loop endless
            l = mid+1
        }
    }
    return r
}

func searchRange(nums []int, target int) []int {
    r := upper_bound(nums, target)
    l := upper_bound(nums, target-1)
    if l == r {
        return []int{-1, -1}
    } else {
        return []int{l, r-1}
    }
}
