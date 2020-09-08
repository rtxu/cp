func destCity(paths [][]string) string {
    n := len(paths)
    outdegree := make(map[string]int)
    for i := 0; i < n; i++ {
        from, to := paths[i][0], paths[i][1]
        outdegree[to] = outdegree[to]
        outdegree[from]++
    }
    for city, degree := range outdegree {
        if degree == 0 {
            return city
        }
    }
    return ""
}
