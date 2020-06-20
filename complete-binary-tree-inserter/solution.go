/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type CBTInserter struct {
    array []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
    array := make([]*TreeNode, 2)
    array[1] = root
    for i := 1; i < len(array); i++ {
        if array[i].Left != nil {
            array = append(array, array[i].Left)
        }
        if array[i].Right != nil {
            array = append(array, array[i].Right)
        }
    }
    return CBTInserter{array}
}


func (this *CBTInserter) Insert(v int) int {
    current := &TreeNode{v, nil, nil}
    this.array = append(this.array, current)
    index := len(this.array)-1
    parentIndex := index/2
    parent := this.array[parentIndex]
    if index == parentIndex*2 {
        parent.Left = current
    } else {
        parent.Right = current
    }
    return parent.Val
}


func (this *CBTInserter) Get_root() *TreeNode {
    return this.array[1]
}


/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(v);
 * param_2 := obj.Get_root();
 */
