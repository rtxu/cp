func next(n int) int {
    sum := 0
    for n > 0 {
        digit := n % 10
        n /= 10
        sum += digit * digit
    }
    return sum
}

func isHappy(n int) bool {
    slow, fast := n, n
    for next(fast) != 1 && next(next(fast)) != 1 {
        slow = next(slow)
        fast = next(next(fast))
        if slow == fast {
            // cycle detected
            return false
        }
    }
    return true
}
