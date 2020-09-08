// Time: O(n)
// Space: O(1)

func duplicateZeros(arr []int)  {
    n := len(arr)
    shiftCount := 0
    for i := n-1; i >= 0; i-- {
        if arr[i] == 0 {
            shiftCount++
        }
    }
    for i := n-1; i >= 0; i-- {
        if i+shiftCount < n {
            arr[i+shiftCount] = arr[i]
        }
        if arr[i] == 0 {
            shiftCount--
            if i+shiftCount < n {
                arr[i+shiftCount] = 0
            }
        }
    }
}
