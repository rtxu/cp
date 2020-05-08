/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
    if root == nil {
        return nil
    }
    
    if root.Left != nil {
        root.Left.Next = root.Right
    }
    
    rightMost := root.Right
    if rightMost == nil {
        rightMost = root.Left
    }
    
    // rightMost.Next == nil 是一个剪枝策略, 避免重复计算
    // 当前的算法可以确保某一个深度的最左节点遍历完成后, 整个一层的链表建立完成
    // 故, 没必要重复计算 next
    if rightMost != nil && rightMost.Next == nil {
        current := root.Next
        for current != nil {
            if current.Left != nil {
                rightMost.Next = current.Left
                rightMost = rightMost.Next
            }
            if current.Right != nil {
                rightMost.Next = current.Right
                rightMost = rightMost.Next
            }
            
            current = current.Next
        }
    }
    
    connect(root.Left)
    connect(root.Right)
    
    return root
}
