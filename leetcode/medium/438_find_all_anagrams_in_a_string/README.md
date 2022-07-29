# [438. find-all-anagrams-in-a-string](https://leetcode.cn/problems/find-all-anagrams-in-a-string/)

## 题目
找到字符串中所有字母异位词 </br>
给定两个字符串s和 p，找到s中所有p的异位词的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。

示例 1：
~~~
输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
~~~

示例 2：
~~~
输入: s = "abab", p = "ab"
输出: [0,1,2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
~~~

## 题目大意

1. 给定字符串p，从字符串s中找出p的任意组合字符串
2. 返回任意组合字符串的起始索引

## 解题思路

**方法一：**
1. 滑动窗口寻找
2. 构建一个need-map，存储p字符串。构建一个windows-map，存储s中出现的p字符串。其中map对应key为字符，value为出现的次数
3. 定义left窗口左边界，right窗口右边界，valid为有效匹配个数，即windows与need对应key的value相等时加1
4. 通过每次right+1操作遍历整个s字符串
5. 遍历s分为前置处理和后置处理
6. 前置处理：当前遍历到的s字符为need中的字符时，将其添加进windows中，并判断`need[curr]==windows[curr]`是否相等，对应操作valid
7. 后置处理：判断窗口大小，即right-left是否等于len(p)，验证valid是否与need大小相等--对应操作结果集，即将left添加到ans中
8. 后置处理：对应操作left，操作windows-map