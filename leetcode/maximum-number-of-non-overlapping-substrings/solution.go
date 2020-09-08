type Node struct {
    cnt, length int
}

type Interval struct {
    begin, end int
}

type ByEnd []Interval

func (a ByEnd) Len() int { return len(a) }
func (a ByEnd) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByEnd) Less(i, j int) bool { 
    return a[i].end < a[j].end
}

func max(a, b Node) Node {
    if a.cnt < b.cnt {
        return b
    } else if a.cnt > b.cnt {
        return a
    } else {
        if a.length < b.length {
            return a
        } else {
            return b
        }
    }
}

func maxNumOfSubstrings(s string) []string {
    // A(i) 表示 [0, ..., i] 子串可构成的最优解 <cnt, len>，相同 cnt 取 len 最短
    // A(i) = max(A(i-1), Node{A(interval(s[i]).begin-1).cnt + 1, A(interval(s[i]).begin-1).length + len(interval(s[i]))}
    var begin, end [26]int
    n := len(s)
    for i := 0; i < n; i++ {
        ch := int(s[i] - 'a')
        if begin[ch] == 0 {
            begin[ch] = i+1
        }
        end[ch] = i+1
    }
    
    intervals := make([]Interval, 0)
    for i := 0; i < 26; i++ {
        if begin[i] == 0 { continue }
        interval := Interval{begin[i], end[i]}
        valid := true
        for j := interval.begin+1; j < interval.end; j++ {
            ch := int(s[j-1] - 'a')
            if begin[ch] < interval.begin {
                valid = false
                break
            }
            if end[ch] > interval.end {
                interval.end = end[ch]
            }
        }
        if valid {
            intervals = append(intervals, interval)
        }
    }
    sort.Sort(ByEnd(intervals))
    // fmt.Println(intervals)
    
    A := make([]Node, n+1)
    for i, j := 1, 0; i <= n; i++ {
        A[i] = A[i-1]
        if i == intervals[j].end {
            A[i] = max(A[i], Node{
                cnt: A[intervals[j].begin-1].cnt + 1,
                length: A[intervals[j].begin-1].length + (intervals[j].end - intervals[j].begin + 1),
            })
            j++
        }
    }
    var result []string
    for i, j := n, len(intervals)-1; i >= 1;  {
        if A[i] != A[i-1] {
            for j >= 0 && i != intervals[j].end { j-- }
            result = append(result, s[intervals[j].begin-1:intervals[j].end])
            i = intervals[j].begin-1
        } else {
            i--
        }
    }
    return result
}
