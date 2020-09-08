/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func _minCameraCover(root *TreeNode, M map[*TreeNode]int, A *[][3]int) int {
    var v [3]int
    
    if root.Left == nil && root.Right == nil {
        // 叶子节点
        v[0] = 0
        v[1] = 1e9 // impossible
        v[2] = 1
    } else if root.Left == nil {
        // 无左子树
        _minCameraCover(root.Right, M, A)
        
        right := M[root.Right]
        sA := (*A)
        v[0] = sA[right][1]
        v[1] = sA[right][2]
        v[2] = 1 + min(sA[right][0], min(sA[right][1], sA[right][2]))
    } else if root.Right == nil {
        _minCameraCover(root.Left, M, A)
        
        left := M[root.Left]
        sA := (*A)
        v[0] = sA[left][1]
        v[1] = sA[left][2]
        v[2] = 1 + min(sA[left][0], min(sA[left][1], sA[left][2]))
    } else {
        _minCameraCover(root.Left, M, A)
        _minCameraCover(root.Right, M, A)
        
        left, right := M[root.Left], M[root.Right]
        sA := (*A)
        v[0] = sA[left][1] + sA[right][1]
        v[1] = min(sA[left][2] + min(sA[right][1], sA[right][2]), sA[right][2] + min(sA[left][1], sA[left][2]))
        v[2] = 1 + min(sA[left][0], min(sA[left][1], sA[left][2])) + min(sA[right][0], min(sA[right][1], sA[right][2]))
    }
    M[root] = len(M)
    (*A) = append(*A, v)
    return min(v[1], v[2])
}

func minCameraCover(root *TreeNode) int {
    M := make(map[*TreeNode]int)
    A := make([][3]int, 0)
    
    // A[i][0] 表示第 i 个 TreeNode 未覆盖自身的情况下，以其为根的子树所需的最小 Camera 数
    // A[i][1] 表示第 i 个 TreeNode 被覆盖但无相机的情况下，以其为根的子树所需的最小 Camera 数
    // A[i][2] 表示第 i 个 TreeNode 被覆盖且有相机的情况下，以其为根的子树所需的最小 Camera 数
    // A[i][3] 表示第 i 个 TreeNode 被覆盖情况下的最优值 = min(A[i][1], A[i][2])
    // A[i][4] 表示第 i 个 TreeNode 上述3者的最小值
    
    return _minCameraCover(root, M, &A)
}
