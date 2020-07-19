// Time: O(n)
// Space: O(n)
// tag: binary-tree-traversal

func countSubTrees(n int, edges [][]int, labels string) []int {
    E := make([][]int, n)
    m := len(edges)
    for i := 0; i < m; i++ {
        n1, n2 := edges[i][0], edges[i][1]
        E[n1] = append(E[n1], n2)
        E[n2] = append(E[n2], n1)
    }
    
    counts := make([]int, n)
    visited := make([]bool, n)
    postOrder(0, labels, E, visited, counts)
    return counts
}

func postOrder(current int, labels string, E [][]int, visited []bool, counts []int) map[byte]int {
    if visited[current] {
        return nil
    }
    visited[current] = true

    charCount := map[byte]int{
        labels[current]: 1,
    }
    m := len(E[current])
    for i := 0; i < m; i++ {
        next := E[current][i]
        childCharCount := postOrder(next, labels, E, visited, counts)
        // fmt.Println("child", next, current, childCharCount)
        for char, cnt := range childCharCount {
            charCount[char] += cnt
        }
    }
    // fmt.Println(current, charCount)
    
    counts[current] = charCount[labels[current]]
    return charCount
}
