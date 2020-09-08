func appendOne(one []int, all *[][]int) {
    oneCopy := make([]int, len(one)) 
    copy(oneCopy, one)
    *all = append(*all, oneCopy)
}

func backtrack(current []int, count map[int]int, nums []int, result *[][]int) {
    if len(current) == len(nums) {
        appendOne(current, result)
        return
    }
    numsCount := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        numsCount[nums[i]]++
        if numsCount[nums[i]] == count[nums[i]] + 1 {
            current = append(current, nums[i])
            count[nums[i]]++
            backtrack(current, count, nums, result)
            count[nums[i]]--
            current = current[0:len(current)-1]
        }
    }
}

func permuteUnique(nums []int) [][]int {
    var result [][]int
    backtrack([]int{}, map[int]int{}, nums, &result)
    return result
}
