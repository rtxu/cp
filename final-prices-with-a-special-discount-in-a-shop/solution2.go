// Time: O(n)
// Space: O(n)

func finalPrices(prices []int) []int {
    n := len(prices)
    result := make([]int, n)
    
    // 维护一个单调递增的栈，对于每一个 i 从栈顶开始遍历 >= prices[i] 的值，该值打折
    stack := make([]int, n)
    stackSize := 0
    for j := 0; j < n; j++ {
        for stackSize > 0 && prices[stack[stackSize-1]] >= prices[j] {
            result[stack[stackSize-1]] -= prices[j]
            stackSize--
        }
        stack[stackSize] = j
        stackSize++
        result[j] = prices[j]
    }
    return result
}
