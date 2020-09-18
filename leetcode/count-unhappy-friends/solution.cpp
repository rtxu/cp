class Solution {
public:
    int unhappyFriends(int n, vector<vector<int>>& preferences, vector<vector<int>>& pairs) {
        unordered_map<int, int> pairMap;
        for (auto& pair: pairs) {
            pairMap[pair[0]] = pair[1];
            pairMap[pair[1]] = pair[0];
        }
        int count = 0;
        for (int x = 0; x < n; x++) {
            int y = pairMap[x];
            bool found = false;
            for (int i = 0; i < (int)preferences[x].size() && preferences[x][i] != y && !found; i++) {
                int u = preferences[x][i];
                int v = pairMap[u];
                for (auto uPrefer: preferences[u]) {
                    if (uPrefer == v) break;
                    if (uPrefer == x) {
                        found = true;
                        break;
                    }
                }
            }
            if (found) {
                count++;
            }
        }
        return count;
    }
};
