func findMaxLength(nums []int) int {
    // diff[i] 表示 nums 0...i 0 的 cnt - 1 的 cnt
    // 则对于任意位置 i，具有相同 diff 的最远位置 j (j > i) 即为以 i 为起点的最长子数组
    count0, count1 := 0, 0
    n := len(nums)
    diff := make([]int, n)
    rightMostDiffPos := make(map[int]int)
    for i := 0; i < n; i++ {
        if nums[i] == 0 {
            count0++
        } else {
            count1++
        }
        diff[i] = count0 - count1
        rightMostDiffPos[diff[i]] = i
    }
    
    var result int
    if rightMost0, exist := rightMostDiffPos[0]; exist {
        result = rightMost0 + 1
    }
    for i := 0; i < n; i++ {
        j := rightMostDiffPos[diff[i]]
        if j > i && j - i > result {
            result = j - i
        }
    }
    return result
}
