
const EPS = 1e-6

var n int
var x, y []float64

func dist(cx, cy float64) float64 {
    var sum float64
    for i := 0; i < n; i++ {
        a := cx - x[i]
        b := cy - y[i]
        sum += math.Sqrt(a*a + b*b)
    }
    return sum
}

func ternarySearchX(cy float64) float64 {
    left, right := float64(0), float64(100)
    for left + EPS < right {
        m1 := left + (right - left)/3
        m2 := m1 + (right - left)/3
        d1, d2 := dist(m1, cy), dist(m2, cy)
        if d1 < d2 {
            right = m2
        } else {
            left = m1
        }
    }
    return (left + right) / 2
}

func getMinDistSum(positions [][]int) float64 {
    n = len(positions)
    x, y = make([]float64, n), make([]float64, n)
    for i := 0; i < n; i++ {
        x[i], y[i] = float64(positions[i][0]), float64(positions[i][1])
    }
    cx, cy := float64(0), float64(0)
    for i := 0; i < 200; i++ {
        cx = ternarySearchX(cy)
        x, y = y, x
        cx, cy = cy, cx
    }
    
    return dist(cx, cy)
}
