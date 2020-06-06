// Time: O(kn), k 为 0 的个数
// Space: O(1)

func duplicateZeros(arr []int)  {
    n := len(arr)
    shift := func(start int) {
        for i := n-1; i >= start; i-- {
            arr[i] = arr[i-1]
        }
    }
    for i := 0; i < n; i++ {
        if arr[i] == 0 {
            shift(i+1)
            if i+1 < n {
                arr[i+1] = 0
            }
            i++
        }
    }
}
