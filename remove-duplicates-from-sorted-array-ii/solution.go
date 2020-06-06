// Time: O(kn), k 为重复次数 > 2 的数字个数
// Space: O(1)

func removeDuplicates(nums []int) int {
    L := len(nums)
    move := func(start, end, cnt int) {
        for i := start; i < end; i++ {
            nums[i-cnt] = nums[i]
        }
    }
    for i := 0; i < L; i++ {
        if i+1 < L && nums[i] == nums[i+1] {
            j := i+2
            for j < L && nums[j] == nums[i] {j++}
            if j - i > 2 {
                move(j, L, j-i-2)
                L -= j-i-2
                i++
            }
        }
    }
    return L
}
