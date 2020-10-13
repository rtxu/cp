class Solution {
    bool valid(const vector<int>& records) {
        int n = records.size();
        for (int i = 0; i < n; i++) {
            int begin = records[i];
            int count = 0;
            for (int j = i; j < n && records[j] - begin <= 60 && count < 3; j++) {
                count++;
            }
            if (count >= 3) {
                return false;
            }
        }
        return true;
    }
public:
    vector<string> alertNames(vector<string>& keyName, vector<string>& keyTime) {
        map<string, vector<int>> userRecords;
        int n = keyName.size();
        for (int i = 0; i < n; i++) {
            istringstream sin(keyTime[i]);
            int hour, min;
            char tmp;
            sin >> hour >> tmp >> min;
            userRecords[keyName[i]].push_back(hour * 60 + min);
        }
        
        vector<string> result;
        for (auto& entry: userRecords) {
            sort(entry.second.begin(), entry.second.end());
            if (!valid(entry.second)) { 
                result.push_back(entry.first);
            }
        }
        return result;
    }
};
