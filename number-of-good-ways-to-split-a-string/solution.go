func numSplits(s string) int {
    n := len(s)
    var head, tail [26]int
    for i := 0; i < n; i++ {
        tail[int(s[i] - 'a')]++
    }
    var count int
    for i := 0; i < n-1; i++ {
        head[int(s[i] - 'a')]++
        tail[int(s[i] - 'a')]--
        var headCnt, tailCnt int
        for j := 0; j < 26; j++ {
            if head[j] > 0 {
                headCnt++
            }
            if tail[j] > 0 {
                tailCnt++
            }
        }
        if headCnt == tailCnt {
            count++
        }
    }
    return count
}
