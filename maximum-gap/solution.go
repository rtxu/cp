func maximumGap(nums []int) int {
    n := len(nums)
    if n < 2 {
        return 0
    }
    
    minNum, maxNum := nums[0], nums[0]
    for i := 1; i < n; i++ {
        if nums[i] < minNum {
            minNum = nums[i]
        }
        if nums[i] > maxNum {
            maxNum = nums[i]
        }
    }
    
    bucketHeight := (maxNum - minNum) / (n-1)
    if bucketHeight == 0 {
        bucketHeight++
    }
    bucketCount := (maxNum - minNum) / bucketHeight + 1
    bucketMin, bucketMax := make([]int, bucketCount), make([]int, bucketCount)
    for i := 0; i < bucketCount; i++ {
        bucketMin[i] = -1
        bucketMax[i] = -1
    }
    for i := 0; i < n; i++ {
        bucket := (nums[i] - minNum) / bucketHeight
        if bucketMin[bucket] == -1 || nums[i] < bucketMin[bucket] {
            bucketMin[bucket] = nums[i]
        }
        if bucketMax[bucket] == -1 || nums[i] > bucketMax[bucket] {
            bucketMax[bucket] = nums[i]
        }
    }
    
    maxGap := bucketMax[0] - bucketMin[0]
    lastBucket := 0
    for currentBucket := 1; currentBucket < bucketCount; currentBucket++ {
        if bucketMax[currentBucket] != -1 {
            newGap := bucketMin[currentBucket] - bucketMax[lastBucket]
            if newGap > maxGap {
                maxGap = newGap
            }
            lastBucket = currentBucket
        }
    }
    
    return maxGap
}
