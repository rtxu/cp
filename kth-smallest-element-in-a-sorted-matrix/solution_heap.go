// Time: O(K*logN)
// Space: O(N)

type Node struct {
    r, c    int
    val     int
}

type Heap []Node

func (h Heap) up(p int) {
    pp := (p-1)>>1
    for p > 0 && h[p].val < h[pp].val {
        h.swap(pp, p)
        p = pp
        pp = (p-1)>>1
    }
}

func (h Heap) down(p int) {
    n := len(h)
    for p < n {
        l := p*2+1
        min := p
        if l < n && h[l].val < h[min].val {
            min = l
        }
        r := l+1
        if r < n && h[r].val < h[min].val {
            min = r
        }
        if min != p {
            h.swap(min, p)
            p = min
        } else {
            break
        }
    }
}

func (h Heap) swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h Heap) Size() int {
    return len(h)
}

func (h *Heap) Push(v Node) {
    *h = append(*h, v)
    h.up(len(*h)-1)
}

func (h *Heap) Pop() Node {
    top := (*h)[0]
    h.swap(0, len(*h)-1)
    *h = (*h)[0:len(*h)-1]
    h.down(0)
    return top
}


func kthSmallest(matrix [][]int, k int) int {
    n := len(matrix)
    h := make(Heap, 0, n)
    for i := 0; i < n; i++ {
        h.Push(Node{i, 0, matrix[i][0]})
    }
    
    var result int
    for k > 0 && h.Size() > 0 {
        // fmt.Println(h)
        currentMinNode := h.Pop()
        if currentMinNode.c+1 < n {
            h.Push(Node{currentMinNode.r, currentMinNode.c+1, matrix[currentMinNode.r][currentMinNode.c+1]})
        }
        k--
        if k == 0 {
            result = currentMinNode.val
        }
    }
    return result
}
