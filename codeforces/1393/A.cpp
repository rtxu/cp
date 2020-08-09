#include <iostream>

using namespace std;

int main() {
    int T, n;
    cin >> T;
    for (int i = 0; i < T; i++) {
        cin >> n;
        cout << n/2 + 1 << endl;
    }
    return 0;
}
