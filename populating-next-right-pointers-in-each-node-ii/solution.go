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
    Q := make([]*Node, 1)
    Q[0] = root
    
    for len(Q) > 0 {
        nextQ := make([]*Node, 0)
        for i := 0; i < len(Q); i++ {
            if Q[i].Left != nil {
                nextQ = append(nextQ, Q[i].Left)
            }
            if Q[i].Right != nil {
                nextQ = append(nextQ, Q[i].Right)
            }
            if i > 0 {
                Q[i-1].Next = Q[i]
            }
        }
        
        Q = nextQ
    }
    
    return root
    
}
