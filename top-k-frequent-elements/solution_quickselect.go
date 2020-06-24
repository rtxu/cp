// Time: O(nlogn), 最坏情况退化为 O(n^2)，依赖于 pivot 的选取
// Space: O(n)

type Node struct {
    freq int
    num  int
}

func kselect(nodes []Node, k int) {
    n := len(nodes)
    pivot := nodes[0].freq
    j := 1
    swap := func(i, j int) {
        nodes[i], nodes[j] = nodes[j], nodes[i]
    }
    for i := 1; i < n; i++ {
        if nodes[i].freq >= pivot {
            swap(j, i)
            j++
        }
    }
    swap(j-1, 0)
    if j == k {
        return
    } else if j < k {
        kselect(nodes[j:], k-j)
    } else {
        kselect(nodes[:j-1], k)
    }
}

func topKFrequent(nums []int, k int) []int {
    count := make(map[int]int)
    n := len(nums)
    for i := 0; i < n; i++ {
        count[nums[i]]++
    }
    
    nodes := make([]Node, 0, n)
    for n, cnt := range count {
        nodes = append(nodes, Node{cnt, n})
    }
    
    kselect(nodes, k)
    
    result := make([]int, k)
    for i := 0; i < k; i++ {
        result[i] = nodes[i].num
    }
    
    return result
}
