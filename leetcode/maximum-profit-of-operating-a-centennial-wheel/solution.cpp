class Solution {
public:
    int minOperationsMaxProfit(vector<int>& customers, int boardingCost, int runningCost) {
        int profit = 0;
        int waiting = 0;
        int n = customers.size();
        int maxProfit = -1;
        int maxRotation = -1;
        for (int i = 0; !(i >= n && waiting == 0); i++) {
            if (i < n) {
                waiting += customers[i];
            }
            int board = min(4, waiting);
            waiting -= board;
            profit += boardingCost * board - runningCost;
            if (profit > maxProfit) {
                maxProfit = profit;
                maxRotation = i + 1;
            }
        }
        return maxRotation;
    }
};
