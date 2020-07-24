func buildArray(target []int, n int) []string {
    var result []string
    for i, j := 1, 0; i <= n && j < len(target); i++ {
        if i != target[j] {
            result = append(result, "Push", "Pop")
        } else {
            result = append(result, "Push")
            j++
        }
    }
    return result
}
