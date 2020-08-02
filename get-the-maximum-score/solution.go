func umax(current *int64, val int64) {
    if val > *current {
        *current = val
    }
}

func maxSum(nums1 []int, nums2 []int) int {
    /*
    n1, n2 := len(nums1), len(nums2)
    best1[i] 表示从 nums1[i] 开始可以获得的最优值
    best2[j] 表示从 nums2[j] 开始可以获得的最优值
    
    则最终答案 = max{best1[0], best2[0]}
    
    best1[n1-1] = nums1[n1-1]
    best2[n2-1] = nums2[n2-1]
    
    best1[i] = max{best1[i+1] + nums1[i], best2[j+1] + nums2[j] if nums1[i] == nums2[j]}
    
    */
    const MOD = 1e9+7
    
    n1, n2 := len(nums1), len(nums2)
    best1, best2 := make([]int64, n1+1), make([]int64, n2+1)
    best1[n1], best2[n2] = int64(0), int64(0)
    var i, j int
    for i, j = n1-1, n2-1; i >= 0; i-- {
        for ; j >= 0 && nums2[j] > nums1[i]; j-- {
            umax(&best2[j], best2[j+1] + int64(nums2[j]))
        }
        umax(&best1[i], best1[i+1] + int64(nums1[i]))
        if j >= 0 && nums2[j] == nums1[i] {
            umax(&best1[i], best2[j+1] + int64(nums1[i]))
            best2[j] = best1[i]
            j--
        }
    }
    for ; j >= 0; j-- {
        umax(&best2[j], best2[j+1] + int64(nums2[j]))
    }
    // fmt.Println(best1)
    // fmt.Println(best2)
    var result int64
    umax(&result, best1[0])
    umax(&result, best2[0])
    return int(result % MOD)
}
