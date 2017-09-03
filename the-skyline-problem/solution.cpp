
#include <iostream>
#include <set>
#include <vector>
#include <algorithm>
#include <utility>

using namespace std;

const int LEFT = 0;
const int RIGHT = 1;

struct EndPoint {
    EndPoint(int _x, int _height, int _type): x(_x), height(_height), type(_type) {}
    int x, height, type; 
};

bool operator < (const EndPoint& lhs, const EndPoint& rhs) {
    return lhs.x < rhs.x;
}

class Solution {
public:
    vector<pair<int, int>> getSkyline(vector<vector<int>>& buildings) {
        int n = buildings.size();
        vector<EndPoint> endpoints;
        for (int i = 0; i < n; i++) {
            endpoints.push_back(EndPoint(buildings[i][0], buildings[i][2], LEFT));
            endpoints.push_back(EndPoint(buildings[i][1], buildings[i][2], RIGHT));
        }
        sort(endpoints.begin(), endpoints.end());
        
        multiset<int> heights;
        int i = 0;
        vector<pair<int, int>> result;
        while (i < 2*n) {
            int beforeHeight = heights.empty()? 0 : (*heights.rbegin());
            int currentX = endpoints[i].x;
            int j = i+1;
            for (; j < 2*n && endpoints[j].x == currentX; j++) {
            }
            for (; i < j; i++) {
                if (endpoints[i].type == LEFT) {
                    heights.insert(endpoints[i].height);
                } else {
                    heights.erase(heights.find(endpoints[i].height));
                }
            }
            
            int afterHeight = heights.empty()? 0 : (*heights.rbegin());
            if (beforeHeight != afterHeight) {
                result.push_back(make_pair(currentX, afterHeight));
            }
        }
        return result;
    }
};


