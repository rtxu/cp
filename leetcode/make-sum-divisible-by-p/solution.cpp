class Solution {
public:
    int minSubarray(vector<int>& nums, int p) {
        int sum = 0;
        int n = nums.size();
        for (int i = 0; i < n; i++) {
            sum += nums[i];
            sum %= p;
        }
        int targetRemainder = sum;
        if (targetRemainder == 0) {
            return 0;
        }
        int prefixSum = 0;
        unordered_map<int, int> lastPos;
        lastPos[0] = -1;
        int minLen = n;
        for (int i = 0; i < n; i++) {
            prefixSum += nums[i];
            prefixSum %= p;
            int need = (prefixSum - targetRemainder + p) % p;
            if (lastPos.count(need) > 0) {
                minLen = min(minLen, i - lastPos[need]);
            }
            lastPos[prefixSum] = i;
        }
        return minLen == n ? -1 : minLen;
    }
};
