func checkIfCanBreak(s1 string, s2 string) bool {
    n := len(s1)
    i1, i2 := make([]int, 0, n), make([]int, 0, n)
    for i := 0; i < n; i++ {
        i1 = append(i1, int(s1[i] - 'a'))
        i2 = append(i2, int(s2[i] - 'a'))
    }
    sort.Ints(i1)
    sort.Ints(i2)
    
    r1, r2 := true, true
    for i := 0; i < n; i++ {
        if i1[i] > i2[i] {
            r1 = false
        }
        if i2[i] > i1[i] {
            r2 = false
        }
    }
    return r1 || r2
}
