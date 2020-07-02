func nextGreaterElement(nums1 []int, nums2 []int) []int {
    // stack[top-1] >= stack[top] < current-value
    n := len(nums2)
    stack := make([]int, 0)
    next := make([]int, n)
    val2PosMap := map[int]int{}
    for i := 0; i < n; i++ {
        val2PosMap[nums2[i]] = i
        next[i] = -1
        for len(stack) > 0 && nums2[stack[len(stack)-1]] < nums2[i] {
            poped := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            next[poped] = nums2[i]
        }
        stack = append(stack, i)
    }
    
    result := make([]int, len(nums1))
    for i := 0; i < len(nums1); i++ {
        result[i] = next[val2PosMap[nums1[i]]]
    }
    return result
}
