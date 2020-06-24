func topKFrequent(nums []int, k int) []int {
    count := make(map[int]int)
    n := len(nums)
    for i := 0; i < n; i++ {
        count[nums[i]]++
    }
    bucket := make([][]int, n+1)
    for n, cnt := range count {
        bucket[cnt] = append(bucket[cnt], n)
    }
    result := make([]int, 0, k)
    for i := n; i >= 1 && len(result) < k; i-- {
        for j := 0; j < len(bucket[i]) && len(result) < k; j++ {
            result = append(result, bucket[i][j])
        }
    }
    return result
}
