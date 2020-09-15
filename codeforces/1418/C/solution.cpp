#include <bits/stdc++.h>

using namespace std;

void solve()
{
    int n;
    cin >> n;
    vector<int> a(n), prefixSum(n + 1);
    for (int i = 0; i < n; i++)
    {
        cin >> a[i];
        prefixSum[i + 1] = prefixSum[i] + a[i];
    }
    // [l, r)
    auto rangeSum = [&prefixSum](int l, int r) {
        return prefixSum[r] - prefixSum[l];
    };
    /*
    player: 0 means you, 1 means your friend
    T(i, player) 表示 player 完成 i 所需要的最小 skip points
    T(i, player) = 
    if player is 0, min{T(i-1, 1), T(i-2, 1)}
    if player is 1, min{T(i-1, 0) + score(i-1, i), T(i-2, 0) + score(i-2, i)}
    初始：T(0, 0) = 0
    则 min{T(n, player)} 为最终答案
    */
    const int INF = (int)1e9;
    vector<vector<int>> T(n + 1, vector<int>(2, INF));
    T[0][0] = 0;
    for (int i = 1; i <= n; i++)
    {
        for (int player = 0; player < 2; player++)
        {
            if (player == 0)
            {
                T[i][player] = min(T[i - 1][1 - player], T[max(i - 2, 0)][1 - player]);
            }
            else
            {
                T[i][player] = min(T[i - 1][1 - player] + rangeSum(i - 1, i),
                                   T[max(i - 2, 0)][1 - player] + rangeSum(max(i - 2, 0), i));
            }
        }
    }
    cout << min(T[n][0], T[n][1]) << endl;
}

int main()
{
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int t;
    cin >> t;
    for (int case_num = 0; case_num < t; case_num++)
    {
        solve();
    }
    return 0;
}
