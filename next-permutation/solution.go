func nextPermutation(nums []int)  {
    n := len(nums)
    var j int
    for j = n-1; j > 0 && nums[j-1] >= nums[j]; j-- {}
    // j == 0 || nums[j-1] < nums[j]
    if j == 0 {
    } else {
        var k int
        for k = n-1; k >= j && nums[k] <= nums[j-1]; k-- {}
        // nums[k] > nums[j-1]
        nums[j-1], nums[k] = nums[k], nums[j-1]
    }
    // reorder [j:n) from descending to ascending
    for i := 0; i < (n-j)/2; i++ {
        nums[j+i], nums[n-1-i] = nums[n-1-i], nums[j+i]
    }
}
