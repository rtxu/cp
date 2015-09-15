
#include <vector>
#include <unordered_map>

using namespace std;

class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        unordered_map<int, int> m;
        vector<int> indexes;
        for (size_t i = 0; i < nums.size(); ++i) {
            if (m.find(nums[i]) == m.end()) {
                // not found result yet
                m[target-nums[i]] = i;
            } else {
                // found result
                indexes.push_back(m[nums[i]]+1);
                indexes.push_back(i+1);
                break;
            }
        }
        return indexes;
    }
};
