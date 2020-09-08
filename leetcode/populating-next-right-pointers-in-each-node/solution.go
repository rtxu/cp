/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

// the solution is the same as 'populating-next-right-pointers-in-each-node-ii'
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
