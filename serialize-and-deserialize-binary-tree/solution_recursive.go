/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    current int
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return "#"
    }
    
    return fmt.Sprintf("%d,%s,%s", root.Val, this.serialize(root.Left), this.serialize(root.Right))
}

func (this *Codec) build(nodes []string) *TreeNode {
    node := nodes[this.current]
    this.current++
    if node == "#" {
        return nil
    } else {
        treeNode := &TreeNode{}
        fmt.Sscanf(node, "%d", &treeNode.Val)
        treeNode.Left = this.build(nodes)
        treeNode.Right = this.build(nodes)
        return treeNode
    }
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    nodes := strings.Split(data, ",")
    this.current = 0
    
    return this.build(nodes)
}


/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */


