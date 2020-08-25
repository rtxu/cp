func mostVisited(n int, rounds []int) []int {
    cnt := make([]int, n)
    for i := 0; i < len(rounds); i++ {
        rounds[i]--
    }
    
    for i, current := 0, rounds[0]; i != len(rounds)-1; current = (current + 1) % n {
        cnt[current]++
        if current == rounds[i+1] {
            i++
        }
    }
    maxV := -1
    var result []int
    for i := 0; i < n; i++ {
        if cnt[i] > maxV {
            maxV = cnt[i]
            result = []int{i+1}
        } else if cnt[i] == maxV {
            result = append(result, i+1)
        }
    }
    return result
}
