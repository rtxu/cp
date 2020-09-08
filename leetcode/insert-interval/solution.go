func insert(intervals [][]int, newInterval []int) [][]int {
    n := len(intervals)
    if n == 0 {
        return [][]int{newInterval}
    }
    
    var result [][]int
    if newInterval[1] < intervals[0][0] {
        result = make([][]int, n+1)
        copy(result, [][]int{newInterval})
        copy(result[1:], intervals)
        return result
    } else if newInterval[0] > intervals[n-1][1] {
        result = make([][]int, n+1)
        copy(result, intervals)
        copy(result[n:], [][]int{newInterval})
        return result
    } else {
        var i int 
        for i = 0; i < n && intervals[i][0] <= newInterval[1]; i++ {
            if intervals[i][1] < newInterval[0] {
                // left
                result = append(result, intervals[i])
            } else {
                if intervals[i][0] < newInterval[0] {
                    newInterval[0] = intervals[i][0]
                }
                if intervals[i][1] > newInterval[1] {
                    newInterval[1] = intervals[i][1]
                }
            }
        }
        result = append(result, newInterval)
        for ; i < n; i++ {
            result = append(result, intervals[i])
        }
        return result
    }
}
