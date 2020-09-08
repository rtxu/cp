type ByDistance [][]int

func (a ByDistance) Len() int { return len(a) }
func (a ByDistance) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByDistance) Less(i, j int) bool { return a[i][0]*a[i][0]+a[i][1]*a[i][1] < a[j][0]*a[j][0]+a[j][1]*a[j][1] }

func kClosest(points [][]int, K int) [][]int {
    sort.Sort(ByDistance(points))
    return points[0:K]
}
