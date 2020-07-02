func dailyTemperatures(T []int) []int {
    // 右手边 > 当前值的最近元素
    // 维护一个单调非递增数组，每当元素出栈时，说明出栈元素碰到了右手边 > 当前值的元素
    n := len(T)
    result := make([]int, n)
    stack := make([]int, 0)
    for i := 0; i < n; i++ {
        for len(stack) > 0 && T[stack[len(stack)-1]] < T[i] {
            poped := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            result[poped] = i-poped
        }
        stack = append(stack, i)
    }
    return result
}
