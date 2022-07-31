# [5. longest-substring-without-repeating-characters](https://leetcode.cn/problems/longest-substring-without-repeating-characters/)

## 题目:无重复字符的最长子串

给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

**示例1：**

~~~
输入: s = "abcabcbb"
输出: 3 
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
~~~

**示例2：**

~~~
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
~~~

**示例3：**

~~~
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。

~~~

## 题目大意

给定字符串，求解该字符串中无重复且连续的最长子串

## 解题思路

**方法一：**
1. 滑动窗口处理
2. 定义一个windows map，存储已遍历的子字符串
3. 定义left-左边界，right-右边界
4. left：只有当遍历字符在windows中重复时进行增加，right-用于遍历整个给定字符串
5. 遍历s的同时，for循环处理windows重复的元素，遇到重复元素将进行缩容，即增加left值，对应left位置的map值减1
6. 对s的每一次遍历进行取left~right值与maxAns结果取最大值
