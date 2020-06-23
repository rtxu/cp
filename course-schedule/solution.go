package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
    indegree := make([]int, numCourses)
    E := make([][]int, numCourses)
    for i := 0; i < len(prerequisites); i++ {
        c, pre := prerequisites[i][0], prerequisites[i][1]
        indegree[c]++
        E[pre] = append(E[pre], c)
    }
    Q := []int{}
    for i := 0; i < numCourses; i++ {
        if indegree[i] == 0 {
            Q = append(Q, i)
        }
    }
    for len(Q) > 0 {
        current := Q[0]
        Q = Q[1:]
        for i := 0; i < len(E[current]); i++ {
            next := E[current][i]
            indegree[next]--
            if indegree[next] == 0 {
                Q = append(Q, next)
            }
        }
    }
    for i := 0; i < numCourses; i++ {
        if indegree[i] > 0 {
            return false
        }
    }
    return true
}

func main() {
    fmt.Println(canFinish(2, [][]int{[]int{1, 0}}))
    fmt.Println(canFinish(2, [][]int{[]int{0, 1}, []int{1, 0}}))
}
