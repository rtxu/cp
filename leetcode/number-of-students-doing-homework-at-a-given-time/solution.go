func busyStudent(startTime []int, endTime []int, queryTime int) int {
    n := len(startTime)
    var count int
    for i := 0; i < n; i++ {
        if queryTime >= startTime[i] && queryTime <= endTime[i] {
            count++
        }
    }
    return count
}
