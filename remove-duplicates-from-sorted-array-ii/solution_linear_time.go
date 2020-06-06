// Time: O(n)
// Space: O(1)

func removeDuplicates(nums []int) int {
    L := len(nums)
    // track the current number appearence count
    count := 0
    // 当前数组下标
    j := 0
    for i := 0; i < L; i++ {
        if i == 0 || nums[i] != nums[i-1] {
            count = 1
        } else {
            count++
        }
        if count > 2 {
            continue
        }
        
        nums[j] = nums[i]
        j++
    }
    return j
}
