// Time: O(n^2*m*l), n = 100, m = 500, l = 20
// Space: O(nm)
func isSubset(left, right map[string]bool) bool {
    if len(left) > len(right) {
        return false
    }
    for s, _ := range left {
        if _, exist := right[s]; !exist {
            return false
        }
    }
    return true
}

func peopleIndexes(favoriteCompanies [][]string) []int {
    result := make([]int, 0)
    n := len(favoriteCompanies)
    F := make([]map[string]bool, n)
    for i := 0; i < n; i++ {
        F[i] = make(map[string]bool)
        for j := 0; j < len(favoriteCompanies[i]); j++ {
            F[i][favoriteCompanies[i][j]] = true
        }
    }
    
    for i := 0; i < n; i++ {
        isOtherSubset := false
        for j := 0; j < n; j++ {
            if i != j {
                if isSubset(F[i], F[j]) {
                    isOtherSubset = true
                    break
                }
            }
        }
        if !isOtherSubset {
            result = append(result, i)
        }
    }
    return result
}
