package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxSum int

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := dfs(root.Left)
	r := dfs(root.Right)

	max := int(math.Max(float64(root.Val),
		math.Max(float64(root.Val+l), float64(root.Val+r))))
	if root.Val+l+r > maxSum {
		maxSum = root.Val + l + r
	}
	if max > maxSum {
		maxSum = max
	}
	return max
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxSum = root.Val
	dfs(root)

	return maxSum
}

func main() {
	var root *TreeNode
	var expected int
	root = &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	expected = 6
	fmt.Printf("expected: %d, actual: %d\n", expected, maxPathSum(root))
	root = &TreeNode{
		Val: -2,
		Left: &TreeNode{
			Val: -1,
		},
		Right: &TreeNode{
			Val: -3,
		},
	}
	expected = -1
	fmt.Printf("expected: %d, actual: %d\n", expected, maxPathSum(root))
	root = &TreeNode{
		Val: -2,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 0,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
	}
	expected = 8
	fmt.Printf("expected: %d, actual: %d\n", expected, maxPathSum(root))
	root = &TreeNode{
		Val: -2,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
	}
	expected = 9
	fmt.Printf("expected: %d, actual: %d\n", expected, maxPathSum(root))
	root = &TreeNode{
		Val: 9,
		Left: &TreeNode{
			Val: 6,
		},
		Right: &TreeNode{
			Val: -3,
			Left: &TreeNode{
				Val: -6,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: -6,
						Left: &TreeNode{
							Val: -6,
						},
					},
					Right: &TreeNode{
						Val: -6,
					},
				},
			},
		},
	}
	expected = 16
	fmt.Printf("expected: %d, actual: %d\n", expected, maxPathSum(root))
}
