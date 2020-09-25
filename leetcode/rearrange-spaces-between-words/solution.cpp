class Solution {
public:
    string reorderSpaces(string text) {
        istringstream sin(text);
        string word;
        vector<string> words;
        int wordsLen = 0;
        while (sin >> word) {
            words.push_back(word);
            wordsLen += word.size();
        }
        int spaceCnt = text.size() - wordsLen;
        int delimLen = 0, extraLen = spaceCnt;
        if (words.size() > 1) {
            delimLen = spaceCnt / (words.size() - 1);
            extraLen = spaceCnt - (words.size() - 1) * delimLen;
        }
        string result;
        string sep;
        for (auto word: words) {
            result += sep;
            result += word;
            sep = string(delimLen, ' ');
        }
        result += string(extraLen, ' ');
        return result;
    }
};
