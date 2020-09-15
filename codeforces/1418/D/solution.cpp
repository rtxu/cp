#include <bits/stdc++.h>

using namespace std;

int main()
{
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    /*
    分堆位置确定，答案确定
    假设分堆边界为 [p1_left, p1_right], [p2_left, p2_right], p1_right < p2_left
    则答案为 (p1_right - p1_left) + (p2_right - p2_left) = [min, max] - [p1_right, p2_left] 
        = (max - min + 1) - (p2_left - p1_right + 1)
        = max - min - max_gap
    还需要维护 p 的个数，如果 p <= 2，则 cost 为 0
    */

    int n, q;
    cin >> n >> q;
    set<int> piles;
    multiset<int> gaps;
    auto add = [&](int pile) {
        auto it = piles.insert(pile).first;
        if (next(it) != piles.end())
        {
            gaps.insert(*next(it) - *it);
        }
        if (it != piles.begin())
        {
            gaps.insert(*it - *prev(it));
        }
        if (next(it) != piles.end() && it != piles.begin())
        {
            gaps.erase(gaps.find(*next(it) - *prev(it)));
        }
    };
    auto remove = [&](int pile) {
        auto it = piles.find(pile);
        if (next(it) != piles.end())
        {
            gaps.erase(gaps.find(*next(it) - *it));
        }
        if (it != piles.begin())
        {
            gaps.erase(gaps.find(*it - *prev(it)));
        }
        if (next(it) != piles.end() && it != piles.begin())
        {
            gaps.insert(*next(it) - *prev(it));
        }
        piles.erase(it);
    };
    auto query = [&]() {
        if (piles.size() <= 2)
        {
            return 0;
        }
        else
        {
            return *piles.rbegin() - *piles.begin() - *gaps.rbegin();
        }
    };
    for (int i = 0; i < n; i++)
    {
        int pile;
        cin >> pile;
        add(pile);
    }

    cout << query() << '\n';
    while (q--)
    {
        int op, pile;
        cin >> op >> pile;
        if (op == 0)
            remove(pile);
        else
            add(pile);

        cout << query() << '\n';
    }

    return 0;
}
