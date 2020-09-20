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
    int64_t N;
    int X, M;
    cin >> N >> X >> M;

    vector<int> v;
    unordered_map<int, int> index;
    v.push_back(X);
    index[X] = 0;
    int cycleStart = -1, cycleLen = -1;
    int64_t cycleSum = 0;
    for (int i = 1;; i++)
    {
        int next = (1LL * v[i - 1] * v[i - 1]) % M;
        auto it = index.find(next);
        if (it == index.end())
        {
            v.push_back(next);
            index[next] = i;
        }
        else
        {
            cycleStart = index[next];
            cycleLen = i - cycleStart;
            for (int j = cycleStart; j < i; j++)
                cycleSum += v[j];
            break;
        }
    }
    dbg(v);
    int64_t sum = 0;
    for (int i = 0; i < min((int64_t)cycleStart, N); i++)
    {
        sum += v[i];
    }
    if (N >= cycleStart)
    {
        int64_t left = N - cycleStart;
        sum += (left / cycleLen) * cycleSum;
        for (int i = 0; i < left % cycleLen; i++)
        {
            sum += v[cycleStart + i];
        }
    }
    cout << sum << endl;

    return 0;
}
