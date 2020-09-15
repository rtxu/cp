#include <bits/stdc++.h>

using namespace std;

int64_t ceilDiv(int64_t a, int64_t b)
{
    return a / b + (a % b != 0);
}

void solveOneCase()
{
    int64_t x, y, k;
    cin >> x >> y >> k;
    cout << ceilDiv((1LL + y) * k - 1, x - 1) + k << endl;
}

int main()
{
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int testcases;
    cin >> testcases;
    for (int case_num = 0; case_num < testcases; case_num++)
    {
        solveOneCase();
    }
    return 0;
}
