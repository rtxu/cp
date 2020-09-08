func isPathCrossing(path string) bool {
    visited := map[[2]int]bool{}
    current := [2]int{0, 0}
    for i := 0; i < len(path); i++ {
        visited[current] = true
        next := [2]int{}
        switch path[i] {
        case 'N':
            next[0], next[1] = current[0]-1, current[1]
        case 'S':
            next[0], next[1] = current[0]+1, current[1]
        case 'E':
            next[0], next[1] = current[0], current[1]+1
        case 'W':
            next[0], next[1] = current[0], current[1]-1
        }
        if visited[next] {
            return true
        }
        current = next
    }
    return false
}
