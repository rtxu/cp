class Solution {
    struct Border {
        int left = 64;
        int right = -1;
        int up = 64;
        int down = -1;
    };
    struct Vertex {
        set<int> adj;
        int indeg = 0;
    };
public:
    bool isPrintable(vector<vector<int>>& targetGrid) {
        unordered_map<int, Border> borders;
        int n = targetGrid.size(), m = targetGrid[0].size();
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                int c = targetGrid[i][j];
                Border& b = borders[c];
                b.left = min(b.left, j);
                b.right = max(b.right, j);
                b.up = min(b.up, i);
                b.down = max(b.down, i);
            }
        }
        const int MAXC = 64;
        vector<Vertex> v(MAXC);
        auto addDep = [&](int c1, int c2) {
            // c1 must paint first
            auto it = v[c1].adj.insert(c2);
            if (it.second) {
                v[c2].indeg++;
            }
        };
        for (const auto& c: borders) {
            const auto& b = c.second;
            for (int i = b.up; i <= b.down; i++) {
                for (int j = b.left; j <= b.right; j++) {
                    int target = targetGrid[i][j];
                    if (target != c.first) {
                        addDep(c.first, target);
                    }
                }
            }
        }
        int cnt = 0;
        queue<int> q;
        for (int i = 0; i < MAXC; i++) {
            if (v[i].indeg == 0) {
                q.push(i);
                cnt++;
            }
        }
        while (!q.empty()) {
            int current = q.front();
            q.pop();
            for (auto next: v[current].adj) {
                v[next].indeg--;
                if (v[next].indeg == 0) {
                    q.push(next);
                    cnt++;
                }
            }
        }
        
        return cnt == MAXC;
    }
};

