// Time: O(2^len(s3)) 穷举所有可能
// Space: O(1)

func backtrack(s1, s2, s3 string, i1, i2, i3 int) bool {
    if i3 == len(s3) {
        return true
    }
    if i1 < len(s1) && s3[i3] == s1[i1] {
        if backtrack(s1, s2, s3, i1+1, i2, i3+1) {
            return true
        }
    }
    if i2 < len(s2) && s3[i3] == s2[i2] {
        if backtrack(s1, s2, s3, i1, i2+1, i3+1) {
            return true
        }
    }
    return false
}


func isInterleave(s1 string, s2 string, s3 string) bool {
    if len(s1) + len(s2) != len(s3) {
        return false
    }
    
    return backtrack(s1, s2, s3, 0, 0, 0)
}
