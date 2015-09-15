
#include <gtest/gtest.h>

#include <vector>
#include <unordered_map>

using namespace std;

class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        unordered_map<int, int> m;
        vector<int> indexes;
        for (size_t i = 0; i < nums.size(); ++i) {
            auto it = m.find(nums[i]);
            if (it == m.end()) {
                // not found result yet
                m[target-nums[i]] = i;
            } else {
                // found result
                indexes.push_back(it->second+1);
                indexes.push_back(i+1);
                break;
            }
        }
        return indexes;
    }
};

TEST(TwoSumUT, BasicTest) {
    Solution s;
    vector<int> nums = {3, 2, 4};
    vector<int> results = s.twoSum(nums, 6);

    EXPECT_EQ((size_t)2, results.size());
}
