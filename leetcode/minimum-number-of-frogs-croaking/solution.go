func minNumberOfFrogs(croakOfFrogs string) int {
    index := map[byte]int{
        'c': 0,
        'r': 1,
        'o': 2,
        'a': 3,
        'k': 4,
    }
    n := len(croakOfFrogs)
    if n % 5 != 0 {
        return -1
    }
    var result int
    var count [5]int
    for i := 0; i < n; i++ {
        current := index[croakOfFrogs[i]]
        count[current]++
        if current > 0 && count[current-1] < count[current] {
            return -1
        }
        if count[current] > result {
            result = count[current]
        }
        if current == 4 {
            for i := 0; i < 5; i++ {
                count[i]--
            }
        }
    }
    return result
}
