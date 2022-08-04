# [528. random-pick-with-weight](https://leetcode.cn/problems/random-pick-with-weight/)

## 题目
按权重随机选择
给你一个 下标从 0 开始 的正整数数组w ，其中w[i] 代表第 i 个下标的权重。

请你实现一个函数pickIndex，它可以 随机地 从范围 [0, w.length - 1] 内（含 0 和 w.length - 1）选出并返回一个下标。选取下标 i的 概率 为 w[i] / sum(w) 。

例如，对于 w = [1, 3]，挑选下标 0 的概率为 1 / (1 + 3)= 0.25 （即，25%），而选取下标 1 的概率为 3 / (1 + 3)= 0.75（即，75%）。


示例 1：
~~~
输入：
["Solution","pickIndex"]
[[[1]],[]]
输出：
[null,0]
解释：
Solution solution = new Solution([1]);
solution.pickIndex(); // 返回 0，因为数组中只有一个元素，所以唯一的选择是返回下标 0。
~~~

示例 2：
~~~
输入：
["Solution","pickIndex","pickIndex","pickIndex","pickIndex","pickIndex"]
[[[1,3]],[],[],[],[],[]]
输出：
[null,1,1,1,1,0]
解释：
Solution solution = new Solution([1, 3]);
solution.pickIndex(); // 返回 1，返回下标 1，返回该下标概率为 3/4 。
solution.pickIndex(); // 返回 1
solution.pickIndex(); // 返回 1
solution.pickIndex(); // 返回 1
solution.pickIndex(); // 返回 0，返回下标 0，返回该下标概率为 1/4 。

由于这是一个随机问题，允许多个答案，因此下列输出都可以被认为是正确的:
[null,1,1,1,1,0]

~~~

## 题目大意

给定一个数组，将数组值作为权重，随机返回数组下标，且概率满足权重

## 解题思路

**方法一：**
1. 构建前缀和--利用数组的值
2. 根据前缀和数组长度随机生成整数[1, preSum[n-1]]
3. 二分查找第一个大于等于给定值的下标