class Solution {
public:
    int maxUniqueSplit(string s) {
        int n = s.size();
        int gapCnt = n - 1;
        int result = 0;
        for (int mask = 0; mask < (1 << gapCnt); mask++) {
            bool success = true;
            set<string> m;
            int begin = 0;
            int tempMask = mask | (1 << (n-1));
            // cout << "tempMask" << hex << tempMask << endl;
            for (int j = 0; j < n; j++) {
                if ((1 << j) & tempMask) {
                    auto it = m.insert(s.substr(begin, j - begin + 1));
                    begin = j + 1;
                    if (!it.second) {
                        success = false;
                        break;
                    }
                }
            }
            if (success) {
                /*
                cout << "mask: " << hex << mask << ", m.size(): " << m.size() << endl;
                for (auto s: m) {
                    cout << s << " ";
                }
                cout << endl;
                */
                result = max(result, (int)m.size());
            }
        }
        return result;
    }
};
