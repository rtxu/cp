class Solution {
    const int END = 0;
    const int BEGIN = 1;
    struct Point {
        int id;
        int pos;
        int type;
        Point(int _id, int _pos, int _type): id(_id), pos(_pos), type(_type) {}
    };
public:
    vector<int> busiestServers(int k, vector<int>& arrival, vector<int>& load) {
        set<int> availServers;
        for (int i = 0; i < k; i++) { 
            availServers.insert(i); 
        }
        vector<int> usedCount(k, 0);
        int n = arrival.size();
        vector<Point> points(2*n, Point(-1, -1, -1));
        for (int i = 0; i < n; i++) {
            points[i] = Point(i, arrival[i], BEGIN);
            points[n + i] = Point(i, arrival[i] + load[i], END);
        }
        sort(points.begin(), points.end(), [](const Point& p1, const Point& p2){
            if (p1.pos != p2.pos) {
                return p1.pos < p2.pos;
            } else {
                return p1.type < p2.type;
            }
        });
        
        unordered_map<int, int> point2serverMap;
        int currentServer = 0;
        for (const auto& p: points) {
            if (p.type == BEGIN) {
                auto findIt = availServers.lower_bound(currentServer);
                if (findIt == availServers.end() && availServers.size() > 0) {
                    findIt = availServers.begin();
                }
                if (findIt != availServers.end()) {
                    int usedServerId = *findIt;
                    point2serverMap[p.id] = usedServerId;
                    usedCount[usedServerId]++;
                    availServers.erase(findIt);
                } else {
                    // drop the current req
                }
                
                currentServer++;
                currentServer %= k;
            } else {
                if (point2serverMap.count(p.id)) {
                    int serverId = point2serverMap[p.id];
                    availServers.insert(serverId);
                }
            }
        }
        int maxUsedCount = *max_element(usedCount.begin(), usedCount.end());
        /*
        cout << "maxUsedCount: " << maxUsedCount << endl;
        for (const auto& count: usedCount) {
            cout << count << endl;
        }
        */
        vector<int> result;
        result.reserve(k);
        for (int i = 0; i < k; i++) {
            if (usedCount[i] == maxUsedCount) {
                result.push_back(i);
            }
        }
        return result;
    }
};
