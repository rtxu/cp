func nextGreaterElements(nums []int) []int {
    n := len(nums)
    nums = append(nums, nums...)
    result := make([]int, n)
    for i := 0; i < n; i++ { result[i] = -1 }
    
    stack := make([]int, 0)
    for i := 0; i < 2*n; i++ {
        for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if top < n {
                result[top] = nums[i]
            }
        }
        stack = append(stack, i)
    }
    
    return result
}
