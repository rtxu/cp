class Solution {
public:
    int connectTwoGroups(vector<vector<int>>& cost) {
        /**
        T(i, mask) 表示左集合覆盖到第 i 个元素，右集合覆盖情况为 mask 的最优值
        T(i, mask) = min{
            T(i-1, mask) + min{cost[i][k]}，前 i-1 个就已经完成了 mask 的覆盖
            T(i-1, subMask) + Sum{cost[i][mask - subMask]}，前 i-1 个覆盖了 subMask，由 i 来覆盖 mask - subMask
            注意：subMask 不能为空
        }
        */
        int n = cost.size(), m = cost[0].size();
        vector<int> minCost(n);
        for (int i = 0; i < n; i++) {
            minCost[i] = *min_element(cost[i].begin(), cost[i].end());
            // cout << "minCost[" << i << "] = " << minCost[i] << endl;
        }
        unordered_map<int, int> singleBitMap;
        for (int j = 0; j < m; j++) singleBitMap[1<<j] = j;
        vector<vector<int>> sumCost(n, vector<int>(1 << m, 0));
        for (int i = 0; i < n; i++) {
            for (int j = 1; j < (1 << m); j++) {
                sumCost[i][j] = sumCost[i][j - (j & -j)] + cost[i][singleBitMap[j & -j]];
            }
        }
        const int INF = 1e9;
        vector<vector<int>> T(n, vector<int>(1 << m, INF));
        for (int i = 0; i < n; i++) {
            T[i][0] = 0;
            for (int mask = 1; mask < (1 << m); mask++) {
                if (i == 0) {
                    T[i][mask] = sumCost[i][mask];
                } else {
                    for (int sub = mask; sub > 0; sub = (sub - 1) & mask) {
                        if (sub == mask) {
                            T[i][mask] = min(T[i][mask], T[i-1][mask] + minCost[i]);
                        } else {
                            T[i][mask] = min(T[i][mask], T[i-1][sub] + sumCost[i][mask - sub]);
                        }
                        /*
                        if (i == 2 && mask == (1<<m) - 1) {
                            cout << "sub = " << sub << ", result = " << T[i][mask] << endl;
                        }
                        if (sub == 0) break;
                        */
                    }
                }
            }
            // cout << "T[" << i << "] = " << T[i][(1<<m) - 1] << endl;
        }
        return T[n-1][(1<<m) - 1];
    }
};
