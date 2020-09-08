func reverseWords(s string) string {
    words := strings.Split(s, " ")
    n := len(words)
    j := 0
    for i := 0; i < n; i++ {
        if words[i] != "" {
            words[j] = words[i]
            j++
        }
    }
    for i := 0; i < j/2; i++ {
        words[i], words[j-1-i] = words[j-1-i], words[i]
    }
    return strings.Join(words[:j], " ")
}
