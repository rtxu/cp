package main

import "fmt"
import "strings"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type Codec struct {
    
}

type node struct {
    parentIndex int
    val int
    isLeft bool
    treeNode *TreeNode
}

func (n node) toString() string {
    return fmt.Sprintf("(%d, %d, %t)", n.parentIndex, n.val, n.isLeft)
}

func (n *node) fromString(s string) {
    fmt.Sscanf(s, "(%d, %d, %t)", &n.parentIndex, &n.val, &n.isLeft)
}

const SEPERATOR = '$'

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return ""
    }
    
    Q := make([]node, 1)
    Q[0] = node{-1, root.Val, false, root}
    qHead := 0
    
    var buffer strings.Builder
    for qHead < len(Q) {
        current := Q[qHead]
        if buffer.Len() > 0 {
            buffer.WriteByte(SEPERATOR)
        }
        buffer.WriteString(current.toString())
        
        if current.treeNode.Left != nil {
            t := current.treeNode.Left
            Q = append(Q, node{qHead, t.Val, true, t})
        }
        if current.treeNode.Right != nil {
            t := current.treeNode.Right
            Q = append(Q, node{qHead, t.Val, false, t})
        }
        qHead++
    }
    
    return buffer.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    if data == "" {
        return nil
    }
    nodes := strings.Split(data, string(SEPERATOR)) 
    treeNodes := make([]*TreeNode, len(nodes))
    for i := 0; i < len(nodes); i++ {
        str := nodes[i]
        node := &node{}
        node.fromString(str)
        treeNode := &TreeNode{node.val, nil, nil}
        treeNodes[i] = treeNode
        if node.parentIndex != -1 {
            if node.isLeft {
                treeNodes[node.parentIndex].Left = treeNode
            } else {
                treeNodes[node.parentIndex].Right = treeNode
            }
        }
    }
    fmt.Println(nodes)
    return treeNodes[0]
}

func buildTree() *TreeNode {
    root := &TreeNode{1, 
        &TreeNode{2, nil, nil},
        &TreeNode{3, nil, nil},
    }
    return root
}

func main() {
    codec := &Codec{}
    root := buildTree()
    s := codec.serialize(root)
    fmt.Println("serilaized s: ", s)
    r := codec.deserialize(s)
    rs := codec.serialize(r)
    fmt.Println("serilaized rs: ", rs)
}
