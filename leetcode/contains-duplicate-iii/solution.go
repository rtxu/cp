// Time: O(kn)
// Space: O(1)

func abs(n int) int {
    if n < 0 {
        return -n
    } else {
        return n
    }
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
    n := len(nums)
    for i := 0; i < n; i++ {
        for j := i+1; j <= i+k && j < n; j++ {
            if abs(nums[i] - nums[j]) <= t {
                return true
            }
        }
    }
    return false
}
