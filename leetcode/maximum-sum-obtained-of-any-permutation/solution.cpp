class Solution {
    constexpr static int MOD = 1e9+7;
public:
    int maxSumRangeQuery(vector<int>& nums, vector<vector<int>>& requests) {
        int n = nums.size();
        vector<int> count(n);
        for (auto r : requests) {
            count[r[0]]++;
            if (r[1]+1 < n) count[r[1]+1]--;
        }
        for (int i = 1; i < n; i++) {
            count[i] += count[i-1];
        }
        sort(count.begin(), count.end());
        sort(nums.begin(), nums.end());
        int sum = 0;
        for (int i = 0; i < n; i++) {
            sum += (1LL * count[i] * nums[i]) % MOD;
            sum %= MOD;
        }
        return sum;
    }
};
