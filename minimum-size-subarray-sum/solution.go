// Time: O(n)
// Space: O(1)

func minSubArrayLen(s int, nums []int) int {
    currentSum, currentLen := 0, 0
    bestLen := 0
    for i := 0; i < len(nums); i++ {
        currentSum += nums[i]
        currentLen += 1
        for currentSum >= s {
            if bestLen == 0 || currentLen < bestLen {
                bestLen = currentLen
            }
            currentLen--
            currentSum -= nums[i-currentLen]
        }
    }
    return bestLen
}
