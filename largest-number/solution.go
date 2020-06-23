type ByConcat []string

func (a ByConcat) Len() int { return len(a) }
func (a ByConcat) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByConcat) Less(i, j int) bool { return a[i]+a[j] > a[j]+a[i] }

func largestNumber(nums []int) string {
    n := len(nums)
    if n == 0 {
        return ""
    }
    strs := make([]string, n)
    for i := 0; i < n; i++ {
        strs[i] = fmt.Sprint(nums[i])
    }
    sort.Sort(ByConcat(strs))
    if strs[0] == "0" {
        return "0"
    } else {
        return strings.Join(strs, "")
    }
}
