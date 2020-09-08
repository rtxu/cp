
var F = make(map[int]float64)
func factorial(n int) float64 {
    if v, exist := F[n]; exist {
        return v
    }
    if n == 0 {
        return 1
    }
    F[n] = float64(n) * factorial(n-1)
    return F[n]
}

func perm(balls []int) float64 {
    var sum int
    for _, ball := range balls {
        sum += ball
    }
    answer := factorial(sum)
    for _, ball := range balls {
        answer /= factorial(ball)
    }
    return answer
}

func split(totalBallCount int, balls, a, b []int, step, sa, sb int) float64 {
    if sa > totalBallCount / 2 || sb > totalBallCount / 2 {
        return 0
    }
    if step == len(balls) {
        aColorCount, bColorCount := 0, 0
        for i := 0; i < len(balls); i++ {
            if a[i] > 0 {
                aColorCount++
            }
            if b[i] > 0 {
                bColorCount++
            }
        }
        if aColorCount == bColorCount {
            return perm(a) * perm(b)
        } else {
            return 0
        }
    } else {
        var result float64
        for i := 0; i <= balls[step]; i++ {
            a[step] = i
            b[step] = balls[step] - i
            result += split(totalBallCount, balls, a, b, step+1, sa+i, sb+balls[step]-i)
        }
        a[step], b[step] = 0, 0
        return result
    }
}

func getProbability(balls []int) float64 {
    a := make([]int, len(balls))
    b := make([]int, len(balls))
    
    var sum int
    for _, ball := range balls {
        sum += ball
    }
    
    return split(sum, balls, a, b, 0, 0, 0) / perm(balls)
}
