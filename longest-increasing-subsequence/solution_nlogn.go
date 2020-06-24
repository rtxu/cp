// Time: O(nlogn)
// Space: O(n)

func lower_bound(a []int, target int) int {
    n := len(a)
    if n == 0 {
        return -1
    }
    if a[n-1] < target {
        return -1
    }
    left, right := 0, n-1
    for left < right {
        mid := left + (right-left)/2
        if a[mid] >= target {
            right = mid
        } else {
            left = mid+1
        }
    }
    return right
}

func lengthOfLIS(nums []int) int {
    n := len(nums)
    // 精髓：tail[i] 表示长度为 i+1 的 increasing subsequence 的最小 tail
    // O(n^2) 的方法的 LIS 数组中，仅单纯的编码了递增子序列的长度，LIS[i] 表示以 i 作为起点，可以构成的最长子序列长度
    // 而 tail 数组同时编码了序列长度和序列尾部的数值信息，实现了对有效信息的完美压缩，使得数组本身有序，进而可以二分查找
    tail := make([]int, n)
    tailCnt := 0
    for i := 0; i < n; i++ {
        j := lower_bound(tail[:tailCnt], nums[i])
        if j == -1 {
            tail[tailCnt] = nums[i]
            tailCnt++
        } else {
            tail[j] = nums[i]
        }
    }
    return tailCnt
}
