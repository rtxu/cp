class Solution {
public:
    int maximumRequests(int n, vector<vector<int>>& requests) {
        int r = requests.size();
        int mx = 0;
        
        unordered_map<int, int> singleBitMap;
        for (int i = 0; i < r; i++) singleBitMap[1 << i] = i;
        
        auto check = [&](int mask) {
            vector<int> deg(n);
            for (int i = mask; i > 0; i -= (i & -i)) {
                const vector<int>& req = requests[singleBitMap[i & -i]];
                deg[req[0]]++, deg[req[1]]--;
            }
            for (int i = 0; i < n; i++) {
                if (deg[i] != 0) {
                    return false;
                }
            }
            return true;
        };
        for (int mask = (1 << r) - 1; mask > 0; mask--) {
            int count = 0;
            for (int i = mask; i > 0; i -= (i & -i)) count++;
            if (count <= mx) continue;
            if (check(mask)) {
                mx = count;
            }
        }
        return mx;
    }
};
