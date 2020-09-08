const MOD = 1e9+7

var cache = map[int]int{}

func pow2(n int) int {
    if v, exist := cache[n]; exist {
        return v
    }
    if n == 0 {
        return 1
    }
    cache[n] = pow2(n-1)<<1
    if cache[n] > MOD {
        cache[n] -= MOD
    }
    return cache[n]
}

func numSubseq(nums []int, target int) int {
    // 对 nums 进行排序，找到最大的 i 使得 2*nums[i] <= target，则 0...i 中的数字任意出现均满足要求，方案数：2^(i+1)
    // 因 2*nums[i+1] > target，则以 nums[i+1] 为最大值的序列，均需找到一个最小值，满足 nums[i+1]+minimum <= target
    // 最小值可选范围为 j，使得 j 为最大的下标，满足 nums[i+1]+nums[j] <= target，即 0...j 中任选1个，i+1 必选，其余数字随便出现
    // 令 j = 2, i+1 = 5，则 [0, 5] 必选，1-4随便出现，[1,5] 必选，0/2-4 随意，[2,5] 必选，0/1/3/4 随意
    // 对方案进行去重，令 [1,5] 必选的方案中不能出现 0，[2,5] 必选的方案中不能出现 0/1
    // 则方案数为 2^4 + 2^3 + 2^2 = 2^i + 2^(i-1) + ... + 2^(i-j)
    // 而                2^(i+1) = 2^i + 2^(i-1) + ... + 2^(i-j) + ... + 2^0 + 1
    // 则方案数 = 2^(i+1) - (1 + 2^0 + 2^1 + ... + 2^(i-j-1)) = 2^(i+1) - 2^(i-j)
    
    n := len(nums)
    sort.Ints(nums)
    var i int
    for i = 0; i < n && 2*nums[i] <= target; i++ {}
    if i == 0 {
        return 0
    }
    count := pow2(i) - 1
    
    var j int
    for j = i; i < n; i++ {
        for ; j >= 0 && nums[i] + nums[j] > target; j-- {}
        if j < 0 {
            break
        }
        count += pow2(i) - pow2(i-j-1) + MOD
        count %= MOD
    }
    return count
}
