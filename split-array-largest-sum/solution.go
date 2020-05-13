func isPossible(nums []int, threshold int, m int) bool {
    // count 0...i 已经产生了多少分组
    // sum 正在进行的分组，所形成的的和
    var i, sum, count int
    for i = 0; i < len(nums); i++ {
        if nums[i] > threshold {
            return false
        }
        if sum + nums[i] > threshold {
            count++
            sum = nums[i]
        } else {
            sum += nums[i]
        }
    }
    if sum > 0 {
        count++
    }
    return count <= m
}

func splitArray(nums []int, m int) int {
    sum := 0
    max := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        if nums[i] > max {
            max = nums[i]
        }
    }
    
    // 无论如何做，最终结果都 > left
    // 任意分割数字，最终结果都 <= right
    left, right := max-1, sum
    // 循环不变量 left < right
    for left+1 != right {
        middle := (left+right)/2
        if isPossible(nums, middle, m) {
            right = middle
        } else {
            left = middle
        }
    }
    
    return right
}
