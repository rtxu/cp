// Time: O(n^2)，如果 go 里面有红黑树可以优化到 O(nlogn)，需求：有序的数据结构支持 O(logn) 的二分查找和删除
// Space: O(n)

func avoidFlood(rains []int) []int {
    // 两个指针 + 一个 window
    // window 里面维护着当前 full 的 lake
    // rain 指针指向 next rain 的 day
    // 当碰到 rain 指针指向的 lake 已经在 window 出现时，dry 指针需要从上一次 rain 之后开始找是否有可用的空间用来 dry 掉该 lake
    // 如果 dry 指针 > rain 指针，impossible
    n := len(rains)
    full := map[int]int{}
    result := make([]int, n)
    for i := 0; i < n; i++ {
        // fmt.Println("current", i, "full", full, "lake", rains[i], "dry", dry)
        if rains[i] > 0 {
            result[i] = -1
            lake := rains[i]
            if full[lake] > 0 {
                dry := full[lake]
                for dry < i && rains[dry] != 0 { dry++ }
                if dry < i {
                    result[dry] = lake
                    rains[dry] = -1
                } else {
                    return []int{}
                }
            }
            full[lake] = i+1
        } else {
            result[i] = 1
        }
    }
    return result
}
