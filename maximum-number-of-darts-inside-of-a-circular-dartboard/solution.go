// Time: O(n^3)
// Space: O(1)

func calcCircleCenter(p1, p2 []int, r int) [][]float64 {
    fp1, fp2 := []float64{float64(p1[0]), float64(p1[1])}, []float64{float64(p2[0]), float64(p2[1])}
    // 计算 p[i], p[j] 的中点, m
    // 计算从 p[i] 指向 p[j] 的向量的垂直向量, v
    // m 沿 v 正向和反向各行进 sqrt(r^2-p[i]到m的距离的平方) 即为两个可能的圆心，使得 p[i], p[j] 的圆周上
    m := []float64{(fp1[0]+fp2[0])/2, (fp1[1]+fp2[1])/2}
    
    // distance(p1, p2) > 2*r
    sqrDistance := (p2[0]-p1[0])*(p2[0]-p1[0])+(p2[1]-p1[1])*(p2[1]-p1[1])
    if sqrDistance > 2*r*2*r || sqrDistance == 0 {
        return [][]float64{fp1}
    } else if sqrDistance == 2*r*2*r {
        return [][]float64{m}
    }
    
    distance := math.Sqrt(float64(sqrDistance))
    vp1p2 := []float64{fp2[0]-fp1[0], fp2[1]-fp1[1]}
    orthV := []float64{-vp1p2[1]/distance, vp1p2[0]/distance}
    d := math.Sqrt(float64(r*r) - float64(sqrDistance)/4)
    c1 := []float64{m[0]+d*orthV[0], m[1]+d*orthV[1]}
    c2 := []float64{m[0]-d*orthV[0], m[1]-d*orthV[1]}
    return [][]float64{c1, c2}
}

func inCircle(p []int, c []float64, r int) bool {
    p0, p1 := float64(p[0]), float64(p[1])
    // 1e-6 very important. 浮点数的计算难以避免的会产生非常小的误差，特别是判断相等时
    return (p0-c[0])*(p0-c[0]) + (p1-c[1])*(p1-c[1]) <= float64(r*r) + 1e-6
}

func numPoints(points [][]int, r int) int {
    n := len(points)
    var maxInCirclePoint int
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            centers := calcCircleCenter(points[i], points[j], r)
            for c := 0; c < len(centers); c++ {
                coverCnt := 0
                for k := 0; k < n; k++ {
                    if inCircle(points[k], centers[c], r) {
                        coverCnt++
                    }
                }
                if coverCnt > maxInCirclePoint {
                    maxInCirclePoint = coverCnt
                }
            }
        }
    }
    return maxInCirclePoint
}
