func findRightMostLessEqual(nums []int, target int) int {
    n := len(nums)
    l, r := 0, n-1
    for l < r {
        mid := l + (r-l)>>1 + 1  // Punchline: make mid biased to right, otherwise l = mid will make for-loop endless
        if nums[mid] <= target {
            l = mid
        } else {
            r = mid-1
        }
    }
    if l < n && nums[l] <= target {
        return l
    } else {
        return -1
    }
}

func searchRange(nums []int, target int) []int {
    r := findRightMostLessEqual(nums, target)
    if r == -1 || nums[r] != target {
        return []int{-1, -1}
    }
    l := findRightMostLessEqual(nums, target-1) + 1
    return []int{l, r}
}
