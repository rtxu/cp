type ListNode struct {
    val, pos    int
    prev, next  *ListNode
}

func NewList() *ListNode {
    node := ListNode{}
    node.prev = &node   // 指向当前 List 的最后一个节点
    node.next = &node   // 指向当前 List 的第一个节点
    return &node
}

func (l *ListNode) Head() *ListNode {
    if l.next == l {
        return nil
    } else {
        return l.next
    }
}

func (l *ListNode) Append(node *ListNode) {
    node.next = l.prev.next
    node.prev = l.prev
    node.next.prev = node
    node.prev.next = node
}

func (l *ListNode) Remove(node *ListNode) {
    node.prev.next = node.next
    node.next.prev = node.prev
    node.prev = nil
    node.next = nil
}

func (l *ListNode) dump() {
    for current := l.next; current != l; current = current.next {
        fmt.Printf("%d->", current.val)
    }
    fmt.Println("nil")
}

func totalFruit(tree []int) int {
    var maxWindowSize int
    list := NewList()
    
    listNodeMap := make(map[int]*ListNode)
    windowStart := 0
    for i := 0; i < len(tree); i++ {
        if node, exist := listNodeMap[tree[i]]; exist {
            list.Remove(node)
            list.Append(node)
            node.pos = i
        } else {
            node := &ListNode{tree[i], i, nil, nil}
            list.Append(node)
            listNodeMap[tree[i]] = node
        }

        if len(listNodeMap) > 2 {
            headNode := list.Head()
            list.Remove(headNode)
            delete(listNodeMap, headNode.val)
            windowStart = headNode.pos + 1
        }
        
        if i - windowStart + 1 > maxWindowSize {
            maxWindowSize = i-windowStart+1
        }
    }
    return maxWindowSize
}
