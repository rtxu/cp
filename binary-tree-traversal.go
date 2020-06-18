package main

import "fmt"

type TreeNode struct {
  Val int
  Left, Right *TreeNode
}

func preorder_traversal_recursive(root *TreeNode) []int {
  if root == nil {
    return []int{}
  }
  var result []int
  result = append(result, root.Val)
  result = append(result, preorder_traversal_recursive(root.Left)...)
  result = append(result, preorder_traversal_recursive(root.Right)...)
  return result
}

func preorder_traversal_iterate(root *TreeNode) []int {
  if root == nil {
    return []int{}
  }
  stack := make([]*TreeNode, 0)
  var result []int
  var lastPop *TreeNode
  push := func(n *TreeNode) {
    stack = append(stack, n)
    result = append(result, n.Val)
  }
  top := func() *TreeNode {
    return stack[len(stack)-1]
  }
  pop := func() *TreeNode {
    n := top()
    stack = stack[:len(stack)-1]
    return n
  }
  push(root)
  for len(stack) > 0 {
    current := top()
    if lastPop == current.Left {
      if current.Right != nil {
        push(current.Right)
      } else {
        lastPop = pop()
      }
    } else if lastPop == current.Right {
      lastPop = pop()
    } else {
      if current.Left != nil {
        push(current.Left)
      } else if current.Right != nil {
        push(current.Right)
      } else {
        lastPop = pop()
      }
    }
  }
  return result
}

func inorder_traversal_recursive(root *TreeNode) []int {
  if root == nil {
    return []int{}
  }
  var result []int
  result = append(result, inorder_traversal_recursive(root.Left)...)
  result = append(result, root.Val)
  result = append(result, inorder_traversal_recursive(root.Right)...)
  return result
}

func inorder_traversal_iterate(root *TreeNode) []int {
  if root == nil {
    return []int{}
  }
  stack := make([]*TreeNode, 0)
  var result []int
  var lastPop *TreeNode
  push := func(n *TreeNode) {
    stack = append(stack, n)
  }
  top := func() *TreeNode {
    return stack[len(stack)-1]
  }
  pop := func() *TreeNode {
    n := top()
    stack = stack[:len(stack)-1]
    return n
  }
  push(root)
  for len(stack) > 0 {
    current := top()
    if lastPop == current.Left {
      result = append(result, current.Val)
      if current.Right != nil {
        push(current.Right)
      } else {
        lastPop = pop()
      }
    } else if lastPop == current.Right {
      lastPop = pop()
    } else {
      if current.Left != nil {
        push(current.Left)
      } else if current.Right != nil {
        result = append(result, current.Val)
        push(current.Right)
      } else {
        lastPop = pop()
       	result = append(result, current.Val)
      }
    }
  }
  return result
}

func postorder_traversal_recursive(root *TreeNode) []int {
  if root == nil {
    return []int{}
  }
  var result []int
  result = append(result, postorder_traversal_recursive(root.Left)...)
  result = append(result, postorder_traversal_recursive(root.Right)...)
  result = append(result, root.Val)
  return result
}

func postorder_traversal_iterate(root *TreeNode) []int {
  if root == nil {
    return []int{}
  }
  stack := make([]*TreeNode, 0)
  var result []int
  var lastPop *TreeNode
  push := func(n *TreeNode) {
    stack = append(stack, n)
  }
  top := func() *TreeNode {
    return stack[len(stack)-1]
  }
  pop := func() *TreeNode {
    n := top()
    stack = stack[:len(stack)-1]
    result = append(result, n.Val)
    return n
  }
  push(root)
  for len(stack) > 0 {
    current := top()
    if lastPop == current.Left {
      if current.Right != nil {
        push(current.Right)
      } else {
        lastPop = pop()
      }
    } else if lastPop == current.Right {
      lastPop = pop()
    } else {
      if current.Left != nil {
        push(current.Left)
      } else if current.Right != nil {
        push(current.Right)
      } else {
        lastPop = pop()
      }
    }
  }
  return result
}

func main() {
  /*
       5
     /   \
    2     4
   / \   /
  1   7 3
  */
  root := &TreeNode{
    Val: 5,
    Left: &TreeNode{
      Val: 2,
      Left: &TreeNode{
        Val: 1,
      },
      Right: &TreeNode{
        Val: 7,
      },
    },
    Right: &TreeNode{
      Val: 4,
      Left: &TreeNode{
        Val: 3,
      },
    },
  }
  
  fmt.Println(preorder_traversal_recursive(root))
  fmt.Println(preorder_traversal_iterate(root))
  fmt.Println(inorder_traversal_recursive(root))
  fmt.Println(inorder_traversal_iterate(root))
  fmt.Println(postorder_traversal_recursive(root))
  fmt.Println(postorder_traversal_iterate(root))
}
