func isValid(freq2NumCountMap map[int]int) bool {
    // 成功的模式:
    // 1) 仅包含一个数字 或 所有数字频率都为1
    if len(freq2NumCountMap) == 1 {
        for freq, numCount := range freq2NumCountMap {
            return numCount == 1 || freq == 1
        }
    }
    // 2) 包含多余一个数字且频率不都为1，每个数字出现的次数满足如下模式：
    if len(freq2NumCountMap) == 2 {
        // 2.1) 1 x...x 一个数字仅出现一次，其他数字均出现 x 次，注意：x 可能等于 1
        // 2.2) x+1 x...x
        minFreq := -1
        for freq, _ := range freq2NumCountMap {
            if minFreq == -1 || freq < minFreq {
                minFreq = freq
            }
        }
        // fmt.Println("minFreq", minFreq)
        
        if minFreq == 1 && freq2NumCountMap[minFreq] == 1 {
            return true
        }
        
        return freq2NumCountMap[minFreq+1] == 1
    }
    
    return false
}

func maxEqualFreq(nums []int) int {
    n := len(nums)
    // 统计每一个数字出现的频次
    num2FreqMap := map[int]int{}
    // freq2NumCountMap[i] 表示有多少个数字出现频次为 i
    freq2NumCountMap := map[int]int{}
    
    result := 2
    for i := 0; i < n; i++ {
        oldCount := num2FreqMap[nums[i]]
        num2FreqMap[nums[i]] = oldCount+1

        if oldCount > 0 {
            freq2NumCountMap[oldCount]--
            if freq2NumCountMap[oldCount] == 0 {
                delete(freq2NumCountMap, oldCount)
            }
        }
        freq2NumCountMap[oldCount+1]++
        // fmt.Println(i+1, freq2NumCountMap)
        
        if isValid(freq2NumCountMap) {
            result = i+1
        }
    }
    return result
}
