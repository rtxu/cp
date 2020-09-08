func maxCoins(piles []int) int {
    sort.Ints(piles)
    n := len(piles)
    j := n-1
    var result int
    for i := 0; i < n/3; i++ {
        result += piles[j-1]
        j -= 2
    }
    return result
}
