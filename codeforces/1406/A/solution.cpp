#include <bits/stdc++.h>

using namespace std;

void solve()
{
    int n;
    cin >> n;
    map<int, int> cnt;
    for (int i = 0; i < n; i++)
    {
        int a;
        cin >> a;
        cnt[a]++;
    }
    vector<int> result;
    int k = 2;
    while (k--)
    {
        for (int i = 0;; i++)
        {
            if (cnt[i] > 0)
            {
                cnt[i]--;
            }
            else
            {
                result.push_back(i);
                break;
            }
        }
    }
    cout << result[0] + result[1] << endl;
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
