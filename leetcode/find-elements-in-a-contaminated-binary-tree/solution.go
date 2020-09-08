/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type FindElements struct {
    root *TreeNode
}

func dfs(current *TreeNode, index int) {
    if current == nil {
        return 
    }
    current.Val = index
    dfs(current.Left, index*2+1)
    dfs(current.Right, index*2+2)
}

func Constructor(root *TreeNode) FindElements {
    //dfs(root, 0)
    return FindElements{root}
}

func find(current *TreeNode, label int, path []int, pathIndex int) bool {
    if pathIndex == -1 {
        return true
    }
    toFind := path[pathIndex]
    if label*2+1 == toFind && current.Left != nil {
        return find(current.Left, toFind, path, pathIndex-1)
    }
    if label*2+2 == toFind && current.Right != nil {
        return find(current.Right, toFind, path, pathIndex-1)
    }
    return false
}

// Time: O(depth)
// Space: O(depth)
func (this *FindElements) Find(target int) bool {
    path := make([]int, 0)
    for target != 0 {
        path = append(path, target)
        target = (target-1) >> 1
    }
    return find(this.root, 0, path, len(path)-1)
}


/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
