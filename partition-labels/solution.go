// Time: O(n^2)
// Space: O(n)

func find(x int, f []int) int {
    if x == f[x] {
        return x
    } else {
        f[x] = find(f[x], f)
        return f[x]
    }
}

func partitionLabels(S string) []int {
    lastPos := [26]int{}
    for i := 0; i < 26; i++ {
        lastPos[i] = -1
    }
    n := len(S)
    f := make([]int, n)
    for i := 0; i < n; i++ {
        f[i] = i
    }
    for i := 0; i < n; i++ {
        char := S[i] - 'a'
        if lastPos[char] != -1 {
            for j := i; j > lastPos[char]; j-- {
                f[j] = find(lastPos[char], f)
            }
        } else {
            lastPos[char] = i
        }
    }
    var result []int
    currentGroup, count := 0, 0
    for i := 0; i < n; i++ {
        group := find(i, f)
        // fmt.Println(i, S[i], group)
        if currentGroup != group {
            result = append(result, count)
            currentGroup = group
            count = 1
        } else {
            count++
        }
    }
    result = append(result, count)
    return result
}
