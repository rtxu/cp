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

void solve()
{
    int n, x;
    cin >> n >> x;
    vector<int> a(n);
    int count = 0;
    int sum = 0;
    for (auto &e : a)
    {
        cin >> e;
        count += e == x;
        sum += e;
    }
    if (count == n)
    {
        cout << 0 << endl;
    }
    else
    {
        // 所有人参加比赛，未感染者都变 x，将分数给感染者
        if (count > 0 || (sum % n == 0 && sum / n == x))
        {
            cout << 1 << endl;
        }
        else
        {
            cout << 2 << endl;
        }
    }
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
