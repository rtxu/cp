/*
minh(v) 表示达到价值 v 的最低高度，则显然 h 单调递增
minh(0) = 0
遍历完第一个元素 a[i], h[i] 后，minh[1] = (h[i], a[i])
依次遍历下一个元素 k，在 minh 中找到小于 h[k] 的最近 minh[j].h，新组成 (h[k], minh[j].a + a[k])
已知 h[k] <= minh[j+1].h，
如果 minh[j].a + a[k] > minh[j+1].a，则删掉 j+1
 
需要线段树？红黑树就可以搞定，可以尝试 c++
*/
 
#include <iostream>
#include <set>
 
using namespace std;
typedef long long int64;
 
const int N = 2e5 + 5;
int64 a[N];
int h[N];
 
struct Node {
  int h;
  int64 v;
  bool operator < (const Node& other) const {
    return h < other.h;
  }
  friend ostream& operator<< (ostream &out, const Node& node) { 
    out << "{ " << node.h << ", " << node.v << " }";
    return out;
  }
};
 
int main() {
  int n; 
  cin >> n;
  for (int i = 0; i < n; i++) {
    cin >> h[i];
  }
  for (int i = 0; i < n; i++) {
    cin >> a[i];
  }
 
  set<Node> s;
  s.insert(Node{0, 0});
  for (int i = 0; i < n; i++) {
    auto it = s.lower_bound(Node{h[i], 0});
    auto newV = (--it)->v + a[i];
    it++;
    
    auto end = it;
    for (; end != s.end() && end->v <= newV; end++) ;
    s.erase(it, end);
    s.insert(Node{h[i], newV});
  }
 
  cout << s.rbegin()->v << endl;
  return 0;
}
