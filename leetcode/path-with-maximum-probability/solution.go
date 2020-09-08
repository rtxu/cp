func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
    prob := make([]float64, n)
    prob[start] = 1
    
    e := len(edges)
    for i := 0; i < n-1; i++ {
        for j := 0; j < e; j++ {
            a, b := edges[j][0], edges[j][1]
            // from a to b
            if prob[a] * succProb[j] > prob[b] {
                prob[b] = prob[a] * succProb[j]
            }
            // from b to a
            if prob[b] * succProb[j] > prob[a] {
                prob[a] = prob[b] * succProb[j]
            }
        }
    }
    
    return prob[end]
}
