type Trie struct {
    Children [2]*Trie
}

func (root *Trie) Insert(n int) {
    current := root
    for i := 30; i >= 0; i-- {
        index := 0
        if n & (1<<i) == (1<<i) {
            index = 1
        }
        if current.Children[index] == nil {
            current.Children[index] = &Trie{}
        }
        current = current.Children[index]
    }
}

func dumpTrie(current *Trie, currentPath []int) {
    if current.Children[0] == nil && current.Children[1] == nil {
        fmt.Println(currentPath)
    } else {
        if current.Children[0] != nil {
            currentPath = append(currentPath, 0)
            dumpTrie(current.Children[0], currentPath)
            currentPath = currentPath[:len(currentPath)-1]
        } 
        if current.Children[1] != nil {
            currentPath = append(currentPath, 1)
            dumpTrie(current.Children[1], currentPath)
            currentPath = currentPath[:len(currentPath)-1]
        }
    }
}

func max(num ...int) int {
    if len(num) == 0 {
        panic("zero arg when call `max`")
    }
    v := num[0]
    for i := 1; i < len(num); i++ {
        if num[i] > v {
            v = num[i]
        }
    }
    return v
}

func dfs(left, right *Trie, current int, carry int) int {
    if left == nil || right == nil {
        return current
    }
    current = (current << 1) + carry
    v := make([]int, 0, 4)
    if left.Children[0] != nil && right.Children[1] != nil {
        v = append(v, dfs(left.Children[0], right.Children[1], current, 1))
    }
    if left.Children[1] != nil && right.Children[0] != nil {
        v = append(v, dfs(left.Children[1], right.Children[0], current, 1))
    }
    if len(v) == 0 {
        v = append(v, dfs(left.Children[1], right.Children[1], current, 0))
        v = append(v, dfs(left.Children[0], right.Children[0], current, 0))
    }
    // fmt.Println(current, v)
    return max(v...)
}

func findMaximumXOR(nums []int) int {
    root := &Trie{}
    n := len(nums)
    if n < 2 {
        return 0
    }
    for i := 0; i < n; i++ {
        root.Insert(nums[i])
    }
    // dumpTrie(root, []int{})
    
    current := root
    for current != nil && (current.Children[0] == nil || current.Children[1] == nil) {
        if current.Children[0] != nil {
            current = current.Children[0]
        } else if current.Children[1] != nil {
            current = current.Children[1]
        } else {
            current = nil
        }
    }
    if current == nil {
        return 0
    } else {
        return dfs(current.Children[0], current.Children[1], 0, 1)
    }
    
}
