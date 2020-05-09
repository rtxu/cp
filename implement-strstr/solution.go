func strStr(haystack string, needle string) int {
    if len(needle) == 0 {
        return 0
    }
    
    text, pattern := "_" + haystack, "_" + needle
    
    next := make([]int, len(needle)+1)
    next[1] = 0
    j := 0
    for i := 2; i <= len(needle); i++ {
        for j > 0 && pattern[j+1] != pattern[i] {
            j = next[j]
        }
        if pattern[j+1] == pattern[i] {
            j = j+1
        }
        next[i] = j
    }
    
    j = 0
    for i := 1; i <= len(haystack); i++ {
        for j > 0 && pattern[j+1] != text[i] {
            j = next[j]
        }
        if pattern[j+1] == text[i] {
            j = j+1
        }
        if j == len(needle) {
            return i - j
        }
    }
    return -1
}
