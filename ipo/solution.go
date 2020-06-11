
type project struct {
    profit, capital int
}

// minCapitalHeap
type MinCapitalHeap []project

func (h MinCapitalHeap) Len() int { return len(h) }
func (h MinCapitalHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h MinCapitalHeap) Less(i, j int) bool { return h[i].capital < h[j].capital }
func (h MinCapitalHeap) Top() project { return h[0] }


func (h *MinCapitalHeap) Push(v interface{}) {
    *h = append(*h, v.(project))
}

func (h *MinCapitalHeap) Pop() interface{} {
    old := *h
    tail := old[len(old)-1]
    *h = old[0:len(old)-1]
    return tail
}

type MaxProfitHeap struct {
    MinCapitalHeap
}

func (h MaxProfitHeap) Less(i, j int) bool { return h.MinCapitalHeap[i].profit > h.MinCapitalHeap[j].profit }


func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
    n := len(Profits)
    minCapitalHeap := make(MinCapitalHeap, 0, n)
    for i := 0; i < n; i++ {
        heap.Push(&minCapitalHeap, project{Profits[i], Capital[i]})
    }
    var maxProfitHeap MaxProfitHeap
    maxProfitHeap.MinCapitalHeap = make(MinCapitalHeap, 0, n)
    
    for i := 0; i < k; i++ {
        // 选出所有满足 W 约束的 Project 集合 maxProfitHeap
        for minCapitalHeap.Len() > 0 && minCapitalHeap.Top().capital <= W {
            p := heap.Pop(&minCapitalHeap).(project)
            heap.Push(&maxProfitHeap, p)
        }
        
        // 在 maxProfitHeap 中选择 Profit 最大的
        if maxProfitHeap.Len() > 0 {
            maxProfit := heap.Pop(&maxProfitHeap).(project).profit
            W += maxProfit
        } else {
            break
        }
    }
    
    return W
}
