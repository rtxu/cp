class Solution {
public:
    int sumOddLengthSubarrays(vector<int>& arr) {
        int n = arr.size();
        vector<int> sum(n+1);
        for (int i = 0; i < n; i++) {
            sum[i+1] = sum[i] + arr[i];
        }
        int result = 0;
        for (int l = 1; l <= n; l += 2) {
            for (int i = 0; i + l <= n; i++) {
                result += sum[i+l] - sum[i];
            }
        }
        return result;
    }
};
