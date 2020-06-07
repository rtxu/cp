package main

import "fmt"

type IntHeap struct {
    values  []int
    pos     map[int]map[int]bool
    less    func(i, j int) bool
}

func NewHeap(less func(i, j int) bool) *IntHeap {
    return &IntHeap{
        values: make([]int, 0),
        pos:    make(map[int]map[int]bool),
        less:   less,
    }
}

func (h IntHeap) swap(i, j int) {
    delete(h.pos[h.values[i]], i)
    delete(h.pos[h.values[j]], j)
    h.values[i], h.values[j] = h.values[j], h.values[i]
    h.pos[h.values[i]][i] = true
    h.pos[h.values[j]][j] = true
}

func (h IntHeap) up(p int) {
    pp := (p-1) >> 1
    for p > 0 && h.less(h.values[p], h.values[pp]) {
        h.swap(p, pp)
        p = pp
        pp = (p-1) >> 1
    }
}

func (h IntHeap) down(p int) bool {
    n := len(h.values)
    //fmt.Println("down", p, h.values)
    isMove := false
    for p < n {
        min := p
        l := p*2+1
        if l < n && h.less(h.values[l], h.values[min]) { min = l }
        r := l+1
        if r < n && h.less(h.values[r], h.values[min]) { min = r }
        if min != p {
            h.swap(min, p)
            p = min
            isMove = true
        } else {
            break
        }
    }
    //fmt.Println("down done", h.values)
    return isMove
}

func (h IntHeap) adjust(p int) {
    if !h.down(p) {
        h.up(p)
    }
}

func (h IntHeap) Len() int {
    return len(h.values)
}

func (h *IntHeap) Push(v int) {
    h.values = append(h.values, v)
    if _, exist := h.pos[v]; !exist {
        h.pos[v] = make(map[int]bool)
    }
    h.pos[v][len(h.values)-1] = true
    h.adjust(len(h.values)-1)
}

func (h IntHeap) Top() int {
    return h.values[0]
}

func (h *IntHeap) Pop() int {
    n := len(h.values)
    top := h.values[0]
    h.swap(0, n-1)
    h.removePos(n-1)
    h.values = h.values[0:n-1]
    h.adjust(0)
    return top
}

func (h *IntHeap) removePos(p int) {
    v := h.values[p]
    delete(h.pos[v], p)
    if len(h.pos[v]) == 0 {
        delete(h.pos, v)
    }
}

func (h *IntHeap) RemoveVal(v int) bool {
    pMap, exist := h.pos[v]
    if !exist {
        return false
    }
    var p int
    for p, _ = range pMap {
        break
    }
    n := len(h.values)
    h.swap(p, n-1)
    h.removePos(n-1)
    h.values = h.values[0:n-1]
    if p != n-1 {
        h.adjust(p)
    }
    return true
}

func (h IntHeap) Slice() []int {
    return h.values
}

type MedianFinder struct {
    small, large *IntHeap
}

func NewMedianFinder() *MedianFinder {
    return &MedianFinder{
        small: NewHeap(func(i, j int) bool { return i > j }),
        large: NewHeap(func(i, j int) bool { return i < j }),
    }
}

func (m MedianFinder) Len() int {
    return m.small.Len() + m.large.Len()
}

func (m *MedianFinder) AddNum(v int) {
    m.small.Push(v)
    //fmt.Println("after small push", v, m.small.Slice())
    m.large.Push(m.small.Pop())
    if m.small.Len() < m.large.Len() {
        m.small.Push(m.large.Pop())
    }
}

func (m MedianFinder) GetMedian() float64 {
    //fmt.Println(m.small.Slice(), m.large.Slice())
    if m.Len() % 2 == 1 {
        return float64(m.small.Top())
    } else {
        return float64(m.small.Top() + m.large.Top()) / 2
    }
}

func (m *MedianFinder) RemoveNum(v int) {
    if !m.small.RemoveVal(v) {
        m.large.RemoveVal(v)
    }
    if m.small.Len() < m.large.Len() {
        m.small.Push(m.large.Pop())
    }
    //fmt.Println("after remove", m.small.Slice(), m.large.Slice())
}

func medianSlidingWindow(nums []int, k int) []float64 {
    finder := NewMedianFinder()
    n := len(nums)
    result := make([]float64, n-k+1)
    for i := 0; i < k; i++ {
        finder.AddNum(nums[i])
    }
    result[0] = finder.GetMedian()
    for i, j := k, 1; i < n; i, j = i+1, j+1 {
        finder.RemoveNum(nums[i-k])
        finder.AddNum(nums[i])
        result[j] = finder.GetMedian()
    }
    return result
}

/*
func main() {
    h := NewHeap(func(i, j int) bool { return i < j})
    h.Push(1)
    h.Push(5)
    h.Push(8)
    h.Push(7)
    fmt.Println(h.values)
    fmt.Println(h.pos)

    h.RemoveVal(5)
    fmt.Println(h.values)
    fmt.Println(h.pos)
}
*/

