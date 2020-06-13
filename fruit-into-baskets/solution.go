func totalFruit(tree []int) int {
    var maxWindowSize int
    windowFruitMap := make(map[int]int)
    windowStart := 0
    for i := 0; i < len(tree); i++ {
        windowFruitMap[tree[i]] = i
        if len(windowFruitMap) > 2 {
            minPos := -1
            var minPosFruit int
            for fruit, pos := range windowFruitMap {
                if minPos == -1 || pos < minPos {
                    minPos = pos
                    minPosFruit = fruit
                }
            }
            windowStart = minPos + 1
            delete(windowFruitMap, minPosFruit)
        }
        
        if i - windowStart + 1 > maxWindowSize {
            maxWindowSize = i-windowStart+1
        }
    }
    return maxWindowSize
}
