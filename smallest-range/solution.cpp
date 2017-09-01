
#include <iostream>
#include <set>
#include <vector>

using namespace std;

struct ElementType {
    ElementType(int v, int b, int c): val(v), belong(b), cur(c) {}
    int val;
    int belong;
    int cur;
};

bool operator<(const ElementType& left, const ElementType right) {
    return left.val < right.val;
}

class Solution {
public:
    vector<int> smallestRange(vector<vector<int>>& nums) {
        auto k = nums.size();
        multiset<ElementType> s;
        for (auto i = 0; i < k; i++) {
            if (nums[i].size() > 0) {
                s.insert(ElementType(nums[i][0], i, 0));
            } else {
                return vector<int>{};
            }
        }
        /*
        for (const auto& e : s) {
            cout << "value: " << e.val << ", belong: " << e.belong << ", pos: " << e.cur << endl;
        }
        */
        vector<int> result;
        while (!s.empty()) {
            auto begin = (s.begin())->val;
            auto end = (s.rbegin())->val;
            auto belong = (s.begin())->belong;
            auto cur = (s.begin())->cur;
            
            if (s.size() == k) {
                if (result.empty() || (end - begin) < (result[1] - result[0])) {
                    result = vector<int>{begin, end};
                }
            } else {
                break;
            }
            
            // cout << "begin: " << begin << ", end: " << end << ", belong: " << belong << ", pos: " << cur << endl;
            if (nums[belong].size() > cur+1) {
                s.insert(ElementType(nums[belong][cur+1], belong, cur+1));
            }
            s.erase(s.begin());
            // cout << s.size() << endl;
        }
        return result;
    }
};

int main()
{
   Solution s;
   vector<vector<int> > input;
   vector<int> result;
   
   input = vector<vector<int>>{
       vector<int>{4, 10, 15, 24,26},
       vector<int>{0, 9, 12, 20},
       vector<int>{5, 18, 22, 30},
   };
   result = s.smallestRange(input);
   cout << "range: [";
   for (auto i : result) {
       cout << i << ",";
   }
   cout << "]" << endl;
   
   return 0;
}


