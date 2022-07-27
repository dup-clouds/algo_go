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
2. 定义左边界left、有边界right，遍历right，valid-标识有效匹配个数
3. 循环遍历给定字符串
4. 对应判断是否添加如windowsMap中，valid对应更新
5. 判断left-right边界大小是否等于目标字符串大小
6. 满足5时，判断valid值是否等于needMap大小，等于则说明成立
7. 在6中不成立时，进行left后移，同时判断left位置元素是否为needMap中，同时对valid+windowsMap进行更新