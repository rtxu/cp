// Time: O(n)
// Space: O(1)
// ref: https://leetcode.com/problems/contains-duplicate-iii/discuss/61645/AC-O(N)-solution-in-Java-using-buckets-with-explanation

const MIN_VALUE = math.MinInt32

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
    if t < 0 || k <= 0 {
        return false
    }
    n := len(nums)
    bucketMap := make(map[int64]int64)
    
    for i := 0; i < n; i++ {
        num := int64(nums[i]) - MIN_VALUE
        bucket := num / (int64(t)+1)
        
        if _, exist := bucketMap[bucket]; exist {
            return true
        }
        if less, exist := bucketMap[bucket-1]; exist && num - less <= int64(t) {
            return true
        }
        if greater, exist := bucketMap[bucket+1]; exist && greater - num <= int64(t) {
            return true
        }
        
        if i >= k {
            bucket := (int64(nums[i-k]) - MIN_VALUE) / (int64(t)+1)
            delete(bucketMap, bucket)
        }
        
        bucketMap[bucket] = num
    }
    return false
}
