class Solution {
public:
    int minOperations(vector<string>& logs) {
        stack<string> s;
        for (auto op : logs) {
            if (op == "../") {
                if (s.size() > 0) {
                    s.pop();
                }
            } else if (op == "./") {
                
            } else {
                s.push(op);
            }
        }
        return s.size();
    }
};
