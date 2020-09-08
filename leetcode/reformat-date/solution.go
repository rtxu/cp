var monthMap = map[string]int {
    "Jan": 1,
    "Feb": 2,
    "Mar": 3,
    "Apr": 4,
    "May": 5,
    "Jun": 6,
    "Jul": 7,
    "Aug": 8,
    "Sep": 9,
    "Oct": 10,
    "Nov": 11,
    "Dec": 12,
}

func reformatDate(date string) string {
    fields := strings.Split(date, " ")
    dayStr, monthStr, yearStr := fields[0], fields[1], fields[2]
    day := 0
    for i := 0; i < len(dayStr); i++ {
        if dayStr[i] >= '0' && dayStr[i] <= '9' {
            day *= 10
            day += int(dayStr[i] - '0')
        }
    }
    return fmt.Sprintf("%s-%02d-%02d", yearStr, monthMap[monthStr], day)
}
