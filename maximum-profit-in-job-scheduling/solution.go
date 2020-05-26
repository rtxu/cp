type job struct {
    startTime int
    endTime int
    profit int
}

type ByEndTime []job

func (a ByEndTime) Len() int { return len(a)}
func (a ByEndTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByEndTime) Less(i, j int) bool { return a[i].endTime < a[j].endTime }

func jobScheduling(startTime []int, endTime []int, profit []int) int {
    // 将 endTime 排序，从小到大
    // P[i] 表示截止到 endTime[i] 获得的最大收益
    // P[i] = max(P[i-1], P[startTime_k] + profit[k]) 遍历以 i 作为 endTime 的所有 job k
    
    n := len(startTime)
    jobs := make([]job, n)
    for i := 0; i < n; i++ {
        jobs[i] = job{startTime[i], endTime[i], profit[i]}
    }
    sort.Sort(ByEndTime(jobs))
    
    allTimeMap := make(map[int]bool)
    for i := 0; i < n; i++ {
        allTimeMap[startTime[i]] = true
        allTimeMap[endTime[i]] = true
    }
    timeCnt := len(allTimeMap)
    allTime := make([]int, timeCnt)
    var i int
    for t, _ := range allTimeMap {
        allTime[i] = t
        i++
    }
    sort.Ints(allTime)
    allTimeIndex := make(map[int]int)
    for i := 0; i < timeCnt; i++ {
        allTimeIndex[allTime[i]] = i
    }
    P := make([]int, timeCnt)
    var j int
    for i := 1; i < timeCnt; i++ {
        P[i] = P[i-1]
        for j < n && allTime[i] == jobs[j].endTime {
            newP := P[allTimeIndex[jobs[j].startTime]] + jobs[j].profit
            if newP > P[i] {
                P[i] = newP
            }
            j++
        }
    }
    return P[timeCnt-1]
}
