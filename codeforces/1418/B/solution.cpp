#include <bits/stdc++.h>

using namespace std;

void solve()
{
    /*
    prefix_sum(n) 固定，不管如何移动
    假设 prefix_sum(n) 为负，则 k 永远为 n，随意输出答案
    假设 prefix_sum(n) >= 0，
        如果我们将所有负数前置，则 all prefix_sum except n is negative
        如果我们将所有正数前置，则 all prefix_sum >= 0
    */
    int n;
    cin >> n;
    vector<int> a(n), locked(n);
    vector<int> canMove;
    canMove.reserve(n);
    for (auto &x : a)
    {
        cin >> x;
    }
    for (int i = 0; i < n; i++)
    {
        cin >> locked[i];
        if (!locked[i])
        {
            canMove.push_back(a[i]);
        }
    }
    sort(canMove.begin(), canMove.end(), greater<int>());
    int j = 0;
    for (int i = 0; i < n; i++)
    {
        if (!locked[i])
        {
            a[i] = canMove[j++];
        }
        cout << a[i] << " \n"[i == n - 1];
    }
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
