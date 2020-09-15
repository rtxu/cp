#include <bits/stdc++.h>

using namespace std;

void solve()
{
    int n;
    cin >> n;
    vector<int> neg, pos;
    neg.reserve(n), pos.reserve(n);
    for (int i = 0; i < n; i++)
    {
        int a;
        cin >> a;
        if (a > 0)
        {
            pos.push_back(a);
        }
        else if (a < 0)
        {
            neg.push_back(a);
        }
    }
    int64_t result = INT64_MIN;
    sort(pos.begin(), pos.end());
    sort(neg.begin(), neg.end(), greater<int>());
    if (pos.size() >= 5)
    {
        int i = int(pos.size()) - 5;
        result = max(result, 1LL * pos[i] * pos[i + 1] * pos[i + 2] * pos[i + 3] * pos[i + 4]);
    }
    if (pos.size() >= 3 && neg.size() >= 2)
    {
        int i = int(pos.size()) - 3, j = int(neg.size()) - 2;
        result = max(result, 1LL * pos[i] * pos[i + 1] * pos[i + 2] * neg[j] * neg[j + 1]);
    }
    if (pos.size() >= 1 && neg.size() >= 4)
    {
        int i = int(pos.size()) - 1, j = int(neg.size()) - 4;
        result = max(result, 1LL * pos[i] * neg[j] * neg[j + 1] * neg[j + 2] * neg[j + 3]);
    }
    if (int(pos.size() + neg.size()) < n)
    {
        result = max(result, 0LL);
    }
    if (pos.size() >= 4 && neg.size() >= 1)
    {
        result = max(result, 1LL * pos[0] * pos[1] * pos[2] * pos[3] * neg[0]);
    }
    if (pos.size() >= 2 && neg.size() >= 3)
    {
        result = max(result, 1LL * pos[0] * pos[1] * neg[0] * neg[1] * neg[2]);
    }
    if (neg.size() >= 5)
    {
        result = max(result, 1LL * neg[0] * neg[1] * neg[2] * neg[3] * neg[4]);
    }
    cout << result << endl;
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