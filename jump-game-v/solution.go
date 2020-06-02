func calc(i, d int, arr, A []int) int {
    if A[i] > 0 {
        return A[i]
    }
    
    result := 1
    check := func(j int) {
        A[j] = calc(j, d, arr, A)
        if A[j] + 1 > result {
            result = A[j] + 1
        }
    }
    
    // jump left
    for j := i-1; j >= i-d && j >= 0 && arr[j] < arr[i]; j-- {
        check(j)
    }
    // jump right
    for j := i+1; j <= i+d && j < len(arr) && arr[j] < arr[i]; j++ {
        check(j)
    }
    
    A[i] = result
    //fmt.Println(i, A[i])
    return A[i]
}

func maxJumps(arr []int, d int) int {
    // A[i] 为从 i 起跳可以获得的最大覆盖数
    
    n := len(arr)
    A := make([]int, n)
    var result int
    for i := 0; i < n; i++ {
        A[i] = calc(i, d, arr, A)
        if A[i] > result {
            result = A[i]
        }
    }
    return result
}
