#include <bits/stdc++.h>

using namespace std;

template <typename T>
ostream &operator<<(ostream &os, const vector<T> &v)
{
    os << '{';
    string sep;
    for (const auto &x : v)
        os << sep << x, sep = ", ";
    return os << '}';
}
template <typename T, size_t size>
ostream &operator<<(ostream &os, const array<T, size> &arr)
{
    os << '{';
    string sep;
    for (const auto &x : arr)
        os << sep << x, sep = ", ";
    return os << '}';
}
template <typename A, typename B>
ostream &operator<<(ostream &os, const pair<A, B> &p) { return os << '(' << p.first << ", " << p.second << ')'; }

void dbg_out() { cerr << endl; }
template <typename Head, typename... Tail>
void dbg_out(Head H, Tail... T)
{
    cerr << ' ' << H;
    dbg_out(T...);
}

// #define DEBUG
#ifdef DEBUG
#define dbg(...) cerr << "(" << #__VA_ARGS__ << "):", dbg_out(__VA_ARGS__)
#else
#define dbg(...)
#endif

int64_t f(int64_t n)
{
    return n * (n + 1) / 2;
}

void solve()
{
    int64_t n;
    cin >> n;
    int result = 0;
    for (int i = 1; i < 31 && n > 0; i++)
    {
        int64_t t = f((1 << i) - 1);
        if (t <= n)
        {
            result++;
            n -= t;
        }
    }
    cout << result << endl;
}

int main()
{
    ios::sync_with_stdio(false);
#ifndef DEBUG
    cin.tie(nullptr);
#endif

    int t;
    cin >> t;
    for (int case_num = 0; case_num < t; case_num++)
    {
        solve();
    }
    return 0;
}
