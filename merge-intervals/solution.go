type interval struct {
    begin, end int
}

type byBegin []interval

func (a byBegin) Len() int { return len(a) }
func (a byBegin) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byBegin) Less(i, j int) bool { return a[i].begin < a[j].begin || (a[i].begin == a[j].begin && a[i].end < a[j].end) }


// interval1.end > interval2.begin

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func merge(intervals [][]int) [][]int {
    n := len(intervals)
    if n == 0 {
        return [][]int{}
    }
    I := make([]interval, n)
    for i := 0; i < n; i++ {
        I[i] = interval{intervals[i][0], intervals[i][1]}
    }
    sort.Sort(byBegin(I))
    
    var result [][]int
    currentInterval := I[0]
    for i := 1; i < n; i++ {
        if currentInterval.end >= I[i].begin {
            currentInterval.end = max(currentInterval.end, I[i].end)
        } else {
            result = append(result, []int{currentInterval.begin, currentInterval.end})
            currentInterval = I[i]
        }
    }
    result = append(result, []int{currentInterval.begin, currentInterval.end})
    return result
}
