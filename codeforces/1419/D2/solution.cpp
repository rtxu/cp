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

int main()
{
    ios::sync_with_stdio(false);
#ifndef DEBUG
    cin.tie(nullptr);
#endif

    int n;
    cin >> n;
    vector<int> a(n);
    for (auto &x : a)
    {
        cin >> x;
    }
    sort(a.begin(), a.end());
    vector<int> b(n);
    int j = 0;
    for (int i = 1; i < n; i += 2)
    {
        b[i] = a[j++];
    }
    for (int i = 0; i < n; i += 2)
    {
        b[i] = a[j++];
    }
    int count = 0;
    for (int i = 1; i + 1 < n; i += 2)
    {
        if (b[i - 1] > b[i] && b[i] < b[i + 1])
            count++;
    }
    cout << count << endl;
    for (int i = 0; i < n; i++)
    {
        cout << b[i] << " \n"[i == n - 1];
    }
    return 0;
}
