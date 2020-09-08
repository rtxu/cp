
// ref: https://leetcode.com/problems/circular-array-loop/discuss/307542/C%2B%2B-Too-much-confusion-Right-See-this.

func circularArrayLoop(nums []int) bool {
    n := len(nums)
    
    next := func(i int) int {
        return (i + (nums[i]%n) + n) % n
    }
    
    for i := 0; i < n; i++ {
        slow, fast := i, i
        for nums[slow]*nums[next(fast)] > 0 && nums[slow]*nums[next(next(fast))] > 0 {
            slow = next(slow)
            fast = next(next(fast))
            if slow == fast {
                if slow == next(slow) {
                    // loop length = 1
                    break
                } else {
                    return true
                }
            }
        }
        
        slow = i
        dir := nums[i]
        for dir * nums[slow] > 0 {
            nextSlow := next(slow)
            nums[slow] = 0
            slow = nextSlow
        }
    }
    return false
}
