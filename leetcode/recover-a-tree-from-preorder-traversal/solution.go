/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isDigit(b byte) bool {
  return b >= '0' && b <= '9'
}

type Stack []*TreeNode

func (s Stack) Top() *TreeNode {
  return s[len(s)-1]
}

func (s Stack) Size() int {
  return len(s)
}

func (s Stack) Empty() bool {
  return s.Size() == 0
}

func (s *Stack) Push(v *TreeNode) {
  *s = append(*s, v)
}

func (s *Stack) Pop() *TreeNode {
  old := *s
  *s = old[:len(old)-1]
  return old.Top()
}

func recoverFromPreorder(S string) *TreeNode {
  stack := make(Stack, 0)
  var num int
  var depth int
  for i := 0; i < len(S); i++ {
    if isDigit(S[i]) {
      num = int(S[i]-'0')
      var j int
      for j = i+1; j < len(S) && isDigit(S[j]); j++ {
        num = num * 10 + int(S[j]-'0')
      }
      i = j-1
      
      node := &TreeNode{num, nil, nil}
      if stack.Empty() {
        // root node
        stack.Push(node)
      } else {
        topNodeDepth := stack.Size() - 1
        if depth == topNodeDepth+1 {
          // left child of top node
          stack.Top().Left = node
          stack.Push(node)
        } else if depth == topNodeDepth {
          // right child of the parent of top node
          stack.Pop()
          stack.Top().Right = node
          stack.Push(node)
        } else {
          for depth != topNodeDepth+1 {
            stack.Pop()
            topNodeDepth = stack.Size() - 1
          }
          stack.Top().Right = node
          stack.Push(node)
        }
      }
      depth = 0
    } else {
      depth++
    }
  }
  
  return stack[0]
}
