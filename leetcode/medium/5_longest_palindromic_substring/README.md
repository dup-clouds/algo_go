# [5. longest_palindromic_substring](https://leetcode.cn/problems/longest-palindromic-substring/)

## 题目

给你一个字符串 s，找到 s 中最长的回文子串。

**示例1：**

~~~
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
~~~

**示例2：**

~~~
输入：s = "cbbd"
输出："bb"
~~~

## 题目大意

回文字符串： a ab abba aba
非回文字符串： aab abb
从已有字符串中寻找最长的回文子串

## 解题思路

**方法一：**
1. 回文串具有中心元素特性
2. 奇数回文串 中心点只有一个 aba
3. 偶数回文串 中心点有两个 abba
4. 循环遍历字符串中的每一个元素，向该元素的左右两端扩展(基于奇数或偶数)，返回最大回文串，将所有回文串取最大长度即可
