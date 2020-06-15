type TreeAncestor struct {
    parent []int
}


func Constructor(n int, parent []int) TreeAncestor {
    return TreeAncestor{parent}
}


func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
    for i := 0; i < k && node != -1; i++ {
        node = this.parent[node]
    }
    return node
}


/**
 * Your TreeAncestor object will be instantiated and called as such:
 * obj := Constructor(n, parent);
 * param_1 := obj.GetKthAncestor(node,k);
 */
