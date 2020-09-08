type ByLen []string

func (a ByLen) Len() int { return len(a) }
func (a ByLen) Less(i, j int) bool { return len(a[i]) < len(a[j]) }
func (a ByLen) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func arrangeWords(text string) string {
    word := strings.Split(text, " ")
    sort.Stable(ByLen(word))
    var result strings.Builder
    for i := 0; i < len(word); i++ {
        if i != 0 {
            result.WriteString(" ")
        }
        if i == 0 {
            if word[0][0] >= 'a' && word[0][0] <= 'z' {
                result.WriteByte(word[0][0] - 'a'  + 'A')
                result.WriteString(word[0][1:])
            } else {
                result.WriteString(word[0])
            }
        } else {
            if word[i][0] >= 'A' && word[i][0] <= 'Z' {
                result.WriteByte(word[i][0] - 'A'  + 'a')
                result.WriteString(word[i][1:])
            } else {
                result.WriteString(word[i])
            }
        }
    }
    return result.String()
}
