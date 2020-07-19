func numWaterBottles(numBottles int, numExchange int) int {
    full, empty := numBottles, 0
    var result int
    for full > 0 {
        empty += full
        result += full
        full, empty = empty / numExchange, empty % numExchange
    }
    return result
}
