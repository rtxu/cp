# 从 Decision Making 的角度，谈动态规划问题的状态设计

标签（空格分隔）： 草稿箱

---

题目：[1531. String Compression II](https://leetcode.com/problems/string-compression-ii/)

在周赛中，剩余 52 分钟用于解决该题，但因没有思路而放弃。

## 定义状态：为了能够做出最优决策，我们需要哪些信息？

赛后，自己开始推导状态定义和转移方程
> 方案一：T(i, j, k) 表示 [i, j) 区间删除 k 个字符的最优解

此时其实大脑一片空白，完全由解题惯性驱动，完成上面的定义。

然后开始推导状态转移方程，在推导的过程中发现两个问题：

1. j 一直为 n，不变
2. 问题定义不满足「自包含」属性，导致转移过程中缺少必要的信息

参考：[Remove Boxes 讨论区](https://leetcode.com/problems/remove-boxes/discuss/101310/Java-top-down-and-bottom-up-DP-solutions)，修改状态定义：
> 方案二：T(i, k, L) 表示 [i, n) 区间，左边有 L 个字符与 s[i] 相同，最多删除 k 个字符，的最优解

从这里可以看出，由于 **DP 的状态是当前作出决策的依据，必须充分到可以作出决策**，冗余可以，缺少则不足以解决问题。

## 推导状态转移方程：如何表达决策？

状态转移方程的含义：当前已知的信息为 A，也就是当前所处的状态，我们可以做哪些决策来获得 A 状态下的最优值？

**状态转移方程用来表达作决策的过程**。

本题中，我们能做的决策有两个：

1. 删除字符，if k > 0
2. 保留字符。如果保留，考虑到最终结果中，所有保留字符均以字符组的形式成组压缩，故还需要考虑该字符是作为一个字符组的最后一个字符，还是作为某个字符组的 prefix。

### 删除字符，
```
if i+1 < n && s[i+1] == s[i] {
    T(i+1, k-1, L), k > 0
} else {
    encodeLen(L) + T(i+1, k-1, 0)
}
```
其中 `encodeLen(charCount)` 表示 charCount 个相同字符 encode 后的表达式长度，”aaa" => "a3" => 2

### 保留字符
当前字符作为字符组最后一个字符，`encodeLen(L+1) + T(i+1, k, 0)`
当前字符作为某个字符组的 prefix，找到下一个 j，满足 `j < n && s[j] == s[i] && j-i+1 <= k`，`T(j, k-(j-i-1), L+1)`

时间复杂度：
状态 O(nkn) => O(n^2k)，转移：O(n)，最终 O(n^3k)
其中找到与当前字符相同的下一字符所在位置，可以 O(n) 预处理得到，将复杂度降到 O(n^2k)
空间复杂度：O(nkn) => O(n^2k)
Coding 复杂度：****

## 其他方案（来自讨论区）

### [方案一](https://leetcode.com/problems/string-compression-ii/discuss/757506/Detailed-Explanation-Two-ways-of-DP-from-33-to-100)：四维状态，简单直接的状态转移

本题与 [Remove Boxes](https://leetcode.com/problems/remove-boxes) 的相似性：二者都存在「**通过与后面 component merge 有可能得到更优结果**」的属性。这需要我们在定义状态时，不能单纯的限定范围，还需要表达出**当前进行中还未完成 merge 的状态**，这里称其为**进行中的状态(ongoing)**。

本题中我使用「左边与 s[i] 相同的字符为 L 个」来表达 ongoing 态，有没有其他表达 ongoing 态的方法呢？

周赛排名前列的选手们，普遍采用了如下表达方法：T(i, ch, len, k) 表示 [i, n) 范围内、最多可删除 k 个字符、且左边有 len 个 ch 等待 merge 的情况下的最优解

状态转移：
1. 删除字符。T(i+1, ch, len, k-1)，if k > 0
2. 保留字符
```
if s[i] == ch {
    T(i+1, ch, len+1, k)
} else {
    encodeOne(len) + T(i+1, s[i], 1, k)
}
```

可以看出，这种定义 ongoing 态的方法，将 s[i] 与 ongoing 态解耦，更容易表达题目中的字符分组概念，使得状态转移简单直接，易于实现

时间复杂度：O(n*26*n*k) => O(26n^2k)
空间复杂度：同上
Coding 复杂度：**

### [方案二](https://leetcode.com/problems/string-compression-ii/discuss/757506/Detailed-Explanation-Two-ways-of-DP-from-33-to-100)：无须定义 ongoing 态

这种解法的关键点是：**考虑到最终结果中，所有保留字符均以字符组的形式成组压缩**，那么如果我们可以以成组的方式处理状态，就不会用到 ongoing 态了

状态定义：T(i, k) 表示 [i, n) 范围内、最多可删除 k 个字符的最优解

状态转移：

1. 删除字符。T(i+1, k-1)，if k > 0
2. 保留字符。枚举以 s[i] 为首的字符组的所有可能结束位置 j，encodeOne(sameCount([i, j+1))) + T(j+1, k - diffCount([i, j+1))，if k - diffCount([i, j+1) >= 0

其中，sameCount([i, j+1)) 表示 [i, j+1) 范围内与 s[i] 相同的字符个数，则 diffCount([i, j)) = (j+1-i) - sameCount([i, j+1)) 为 [i, j+1) 内与 s[i] 不同的字符个数，即 [i, j+1) 内被删除的字符个数

时间复杂度：状态 O(nk)，转移O(n)，最终 O(n^2k)
空间复杂度：O(nk)
Coding 复杂度：**

## 最后

本题属于[决策类](https://leetcode.com/discuss/general-discussion/458695/dynamic-programming-patterns) DP 问题，解题时应先分析题目，看都可以做哪些决策，再根据决策设计状态，状态本身是为表达决策服务的。
当发现状态不足、无法做出最优决策时，考虑通过扩充状态的方式为状态增加更多信息。
不同的状态定义，在表达决策时，难度不同，设计不合理的状态会增加表达决策的难度。
