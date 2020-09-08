package main

import "fmt"

const NO_CONSTRAINT = -1

func main() {
    var t int
    fmt.Scanln(&t)
    for ca := 0; ca < t; ca++ {
        var n, k int
		fmt.Scanln(&n, &k)
		var s string
		fmt.Scanln(&s)
		valid := true
		constraints := make([]int, k)
		for i := 0; i < k; i++ {
			constraints[i] = NO_CONSTRAINT
		}
		for i := 0; i < n; i++ {
			if s[i] != '?' {
				c := int(s[i] - '0')
				if constraints[i % k] == NO_CONSTRAINT {
					constraints[i % k] = c
				} else if constraints[i % k] == c {
				} else {
					valid = false
					break
				}
			}
		}
		var cnt [2]int
		for i := 0; i < k; i++ {
			if constraints[i] != NO_CONSTRAINT {
				cnt[constraints[i]]++
			}
		}
		if cnt[0] > k / 2 || cnt[1] > k / 2 {
			valid = false
		}

		if valid {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
    }
}
