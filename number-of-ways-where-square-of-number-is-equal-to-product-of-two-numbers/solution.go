func count(a, b map[int]int) int {
    result := 0
    for v1, cnt1 := range a {
        product := int64(v1) * int64(v1)
        for v2, cnt2 := range b {
            if product % int64(v2) != 0 {
                continue
            }
            v3 := int(product / int64(v2))
            if v2 == v3 {
                if cnt2 >= 2 {
                    result += cnt1 * cnt2 * (cnt2-1) / 2
                }
            } else if v2 < v3 {
                if cnt3, exist := b[v3]; exist {
                    result += cnt1 * cnt2 * cnt3
                }
            } else {
                continue
            }
        }
    }
    return result
}

func numTriplets(nums1 []int, nums2 []int) int {
    cnt1, cnt2 := make(map[int]int), make(map[int]int)
    n1, n2 := len(nums1), len(nums2)
    for i := 0; i < n1; i++ {
        cnt1[nums1[i]]++
    }
    for i := 0; i < n2; i++ {
        cnt2[nums2[i]]++
    }
    return count(cnt1, cnt2) + count(cnt2, cnt1)
}
