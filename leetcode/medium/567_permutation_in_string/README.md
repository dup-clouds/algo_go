# [567. permutation-in-string](https://leetcode.cn/problems/permutation-in-string/)

## 题目
字符串的排列
给你两个字符串s1和s2 ，写一个函数来判断 s2 是否包含 s1的排列。如果是，返回 true ；否则，返回 false 。
换句话说，s1 的排列之一是 s2 的 子串 。


示例 1：
~~~
输入：s1 = "ab" s2 = "eidbaooo"
输出：true
解释：s2 包含 s1 的排列之一 ("ba").
~~~

示例 2：
~~~
输入：s1= "ab" s2 = "eidboaoo"
输出：false
~~~

## 题目大意

判断目标字符串的排列是否在给定字符串中

## 解题思路

**方法一：**
1. 构建目标字符串map-needMap，构建窗口字符串map-windowsMap
2. 定义左边界left、有边界right，遍历right
3. 判断needMap大小是否等于windowsMap
4. 确保windowsMap的连续性，如不存在连续性则移动左边界left