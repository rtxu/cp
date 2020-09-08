func abs(n int) int {
    if n < 0 {
        return -n
    } else {
        return n
    }
}

func getStrongest(arr []int, k int) []int {
    sort.Ints(arr)
    n := len(arr)
    median := arr[(n-1)/2]
    left, right := 0, n-1
    result := make([]int, k)
    for i := 0; i < k; i++ {
        sign := abs(arr[left]-median) - abs(arr[right]-median)
        if sign > 0 {
            result[i] = arr[left]
            left++
        } else if sign < 0 {
            result[i] = arr[right]
            right--
        } else {
            if arr[left] > arr[right] {
                result[i] = arr[left]
                left++
            } else {
                result[i] = arr[right]
                right--
            }
        }
    }
    return result
}
