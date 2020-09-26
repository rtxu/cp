class Solution {
    struct Node {
        int64_t pos = -1;
        int64_t neg = 1;
    };
public:
    int maxProductPath(vector<vector<int>>& grid) {
        int n = grid.size(), m = grid[0].size();
        n++, m++;
        vector<vector<Node>> A(n, vector<Node>(m));
        A[0][1].pos = 1;
        
        int result = -1;
        for (int i = 1; i < n; i++) {
            for (int j = 1; j < m; j++) {
                int v = grid[i-1][j-1];
                if (v == 0) {
                    result = 0;
                } else {
                    int64_t maxPos = max(A[i-1][j].pos, A[i][j-1].pos);
                    int64_t minNeg = min(A[i-1][j].neg, A[i][j-1].neg);
                    Node& current = A[i][j];
                    if (v > 0) {
                        if (maxPos > 0) {
                            current.pos = maxPos * v;
                        }
                        if (minNeg < 0) {
                            current.neg = minNeg * v;
                        }
                    } else {
                        if (maxPos > 0) {
                            current.neg = maxPos * v;
                        }
                        if (minNeg < 0) {
                            current.pos = minNeg * v;
                        }
                    }
                }
                // cout << "(" << A[i][j].pos << ", " << A[i][j].neg << ") ";
            }
            // cout << endl;
        }
        
        const int MOD = 1e9+7;
        if (A[n-1][m-1].pos > 0) {
            result = A[n-1][m-1].pos % MOD;
        }
        return result;
    }
};
