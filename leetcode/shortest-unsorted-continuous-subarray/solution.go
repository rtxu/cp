// Time: O(nlogn)
// Space: O(n)

func findUnsortedSubarray(nums []int) int {
    nums2 := make([]int, len(nums))
    copy(nums2, nums)
    sort.Ints(nums2)
    i, j := 0, len(nums) - 1
    for i < j {
        hasMove := false
        if nums[i] == nums2[i] {
            hasMove = true
            i++
        }
        if nums[j] == nums2[j] {
            hasMove = true
            j--
        }
        if !hasMove {
            break
        }
    }
    if i < j {
        return j+1-i
    } else {
        return 0
    }
}
