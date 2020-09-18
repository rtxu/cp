class Solution {
public:
    int numSpecial(vector<vector<int>>& mat) {
        int nRow = mat.size(), nCol = mat[0].size();
        vector<int> rows(nRow), cols(nCol);
        for (int i = 0; i < nRow; i++) {
            for (int j = 0; j < nCol; j++) {
                rows[i] += mat[i][j];
                cols[j] += mat[i][j];
            }
        }
        int count = 0;
        for (int i = 0; i < nRow; i++) {
            for (int j = 0; j < nCol; j++) {
                if (mat[i][j] == 1 && rows[i] == 1 && cols[j] == 1) {
                    count++;
                }
            }
        }
        return count;
    }
};
