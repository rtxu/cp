// Time: O(n)
// Space: O(1)
func productExceptSelf(nums []int) []int {
    n := len(nums)
    result := make([]int, n)
    for i, prod := 0, 1; i < n; i, prod = i+1, prod*nums[i] {
        result[i] = prod
    }
    for i, prod := n-1, 1; i >= 0; i, prod = i-1, prod*nums[i] {
        result[i] *= prod
    }
    return result
}
