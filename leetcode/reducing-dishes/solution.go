func maxSatisfaction(satisfaction []int) int {
    sort.Ints(satisfaction)
    
    maxSum := 0
    for i := 0; i <= len(satisfaction); i++ {
        sum := 0
        for j := i; j < len(satisfaction); j++ {
            sum += satisfaction[j] * (j-i+1)
        }
        if sum > maxSum {
            maxSum = sum
        }
    }
    
    return maxSum
}
