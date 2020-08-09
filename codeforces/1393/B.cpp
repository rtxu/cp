#include <bits/stdc++.h>
#include <unordered_map>

using namespace std;

unordered_map<int, set<pair<int, int> >::iterator> lengthMap;
set<pair<int, int> > cntSet;

void inc(int L) {
    auto it = lengthMap.find(L);
    int cnt;
    if (it == lengthMap.end()) {
        cnt = 1;
    } else {
        cnt = it->second->first + 1;
        cntSet.erase(it->second);
    }
    auto ret = cntSet.insert(pair<int, int>(cnt, L));
    lengthMap[L] = ret.first;
}

void dec(int L) {
    auto it = lengthMap.find(L);
    int cnt = it->second->first - 1;
    cntSet.erase(it->second);
    auto ret = cntSet.insert(pair<int, int>(cnt, L));
    lengthMap[L] = ret.first;
}

int countAtMost2(int count, int upperLimit) {
    int cnt = 0;
    auto it = cntSet.lower_bound(pair<int, int>(count, 0));
    for (; it != cntSet.end() && it->first < upperLimit && cnt < 2; it++) {
        cnt++;
    }
    // cout << "length: " << count << ", count: " << cnt << endl;
    return cnt;
}

bool valid() {
    const int N = 1e5+5;

    if (countAtMost2(8, N) > 0) {
        return true;
    }
    if (countAtMost2(4, N) >= 2) {
        return true;
    }
    if (countAtMost2(6, N) > 0 && countAtMost2(2, 6) > 0) {
        return true;
    }
    if (countAtMost2(4, N) > 0 && countAtMost2(2, 4) >= 2) {
        return true;
    }

    return false;
}

int main() {
    cout << __cplusplus << endl;
    cout << sizeof(double) << endl;
    cout << sizeof(long double) << endl;
    int n;
    cin >> n;
    for (int i = 0; i < n; i++) {
        int L;
        cin >> L;
        inc(L);
    }
    int q;
    cin >> q;
    for (int i = 0; i < q; i++) {
        string op;
        int L;
        cin >> op >> L;
        if (op == "+") {
            inc(L);
        } else {
            dec(L);
        }

        if (valid()) {
            cout << "YES" << endl;
        } else {
            cout << "NO" << endl;
        }
    }
    return 0;
}
