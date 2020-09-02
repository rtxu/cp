func kidsWithCandies(candies []int, extraCandies int) []bool {
    n := len(candies)
    max := 0
    for i := 0; i < n; i++ {
        if candies[i] > max {
            max = candies[i]
        }
    }
    var result []bool
    for i := 0; i < n; i++ {
        result = append(result, candies[i] + extraCandies >= max)
    }
    return result
}
