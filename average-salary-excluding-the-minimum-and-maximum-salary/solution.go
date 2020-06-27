func average(salary []int) float64 {
    sum := 0
    min := math.MaxInt32
    max := math.MinInt32
    for i := 0; i < len(salary); i++ {
        sum += salary[i]
        if salary[i] < min {
            min = salary[i]
        }
        if salary[i] > max {
            max = salary[i]
        }
    }
    return float64(sum - min - max) / float64(len(salary)-2)
}
