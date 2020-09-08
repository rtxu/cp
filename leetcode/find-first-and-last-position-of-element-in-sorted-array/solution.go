func searchRange(nums []int, target int) []int {
    n := len(nums)
    left, right := 0, n-1
    for left <= right {
        mid := (left+right)>>1
        if nums[mid] < target {
            left = mid+1
        } else if nums[mid] > target {
            right = mid-1
        } else {
            min, max := mid, mid
            if nums[left] == target {
                min = left
            } else {
                for left+1 != min {
                    mid := (left+min)>>1
                    if nums[mid] >= target {
                        min = mid
                    } else {
                        left = mid
                    }
                }    
            }
            if nums[right] == target {
                max = right
            } else {
                for max+1 != right {
                    mid := (max+right)>>1
                    if nums[mid] > target {
                        right = mid
                    } else {
                        max = mid
                    }
                }
            }
            return []int{min, max}
        }
    }
    return []int{-1, -1}
}
