func resemble(a []int) int {
    result := 0
    n := len(a)
    for i := n-1; i >= 0; i-- {
        result = result * 10 + a[i]
    }
    return result
}

func change(x, y int, a []int) {
    n := len(a)
    for i := 0; i < n; i++ {
        if a[i] == x {
            a[i] = y
        }
    }
}

func maxDiff(num int) int {
    big, small := make([]int, 0), make([]int, 0)
    for num > 0 {
        digit := num % 10
        num /= 10
        big = append(big, digit)
        small = append(small, digit)
    }
    n := len(big)
    var i int
    for i = n-1; i >= 0 && big[i] == 9; i-- {}
    x := 9
    if i >= 0 {
        x = big[i]
    }
    change(x, 9, big)

    if small[n-1] != 1 {
        x = small[n-1]
        change(x, 1, small)
    } else {
        for i = n-2; i >= 0 && small[i] <= 1; i-- {}
        x = 0
        if i >= 0 {
            x = small[i]
        }
        change(x, 0, small[0:n-1])
    }
    // fmt.Println("big", big, resemble(big))
    // fmt.Println("sma", small, resemble(small))
    return resemble(big) - resemble(small)
}
