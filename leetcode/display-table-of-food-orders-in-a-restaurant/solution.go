func displayTable(orders [][]string) [][]string {
    foodMap := make(map[string]bool)
    tableMap := make(map[int]map[string]int)
    
    n := len(orders)
    for i := 0; i < n; i++ {
        tableId, _ := strconv.Atoi(orders[i][1])
        food := orders[i][2]
        if _, exist := tableMap[tableId]; !exist {
            tableMap[tableId] = make(map[string]int)
        }
        tableMap[tableId][food]++
        foodMap[food] = true
    }
    
    foods := make([]string, 0, len(foodMap))
    for food, _ := range foodMap {
        foods = append(foods, food)
    }
    sort.Strings(foods)
    tables := make([]int, 0, len(tableMap))
    for tableId, _ := range tableMap {
        tables = append(tables, tableId)
    }
    sort.Ints(tables)
    
    var result [][]string
    var header []string
    header = append(header, "Table")
    header = append(header, foods...)
    result = append(result, header)
    for i := 0; i < len(tables); i++ {
        var row []string
        tableId := tables[i]
        row = append(row, strconv.FormatInt(int64(tableId), 10))
        table := tableMap[tableId]
        for j := 0; j < len(foods); j++ {
            row = append(row, strconv.FormatInt(int64(table[foods[j]]), 10))
        }
        result = append(result, row)
    }
    return result
}
