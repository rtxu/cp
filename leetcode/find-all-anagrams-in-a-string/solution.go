func findAnagrams(s string, p string) []int {
    if len(s) < len(p) {
        return []int{}
    }
    M := map[byte]int{}
    for i := 0; i < len(p); i++ {
        M[p[i]]++
    }
    counter := len(p)
    var result []int
    for i := 0; i < len(s); i++ {
        if M[s[i]] > 0 {
            counter--
        }
        M[s[i]]--
        if i >= len(p) {
            M[s[i-len(p)]]++
            if M[s[i-len(p)]] > 0 {
                counter++
            }
        }
        if counter == 0 {
            result = append(result, i+1-len(p))
        }
    }
    return result
}
