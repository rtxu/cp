class Solution {
    inline int dist(const vector<int>& p1, const vector<int>& p2) {
        return abs(p1[0] - p2[0]) + abs(p1[1] - p2[1]);
    }
public:
    int minCostConnectPoints(vector<vector<int>>& ps) {
        int n = ps.size();
        const int INF = 1e9;
        vector<int> d(n, INF);
        vector<bool> visited(n, false);
        int current = 0;
        d[current] = 0;
        int minCost = 0;
        for (int i = 0; i < n; i++) {
            visited[current] = true;
            minCost += d[current];
            int next = -1;
            for (int j = 0; j < n; j++) if (!visited[j]) {
                d[j] = min(d[j], dist(ps[current], ps[j]));
                if (next == -1 || d[j] < d[next]) {
                    next = j;
                }
            }
            current = next;
        }
        return minCost;
    }
};
