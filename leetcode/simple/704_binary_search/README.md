# [704. binary-search](https://leetcode.cn/problems/binary-search/)

## 题目
二分查找
给定一个n个元素有序的（升序）整型数组nums 和一个目标值target，写一个函数搜索nums中的 target，如果目标值存在返回下标，否则返回 -1。

示例 1：

~~~
输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4
~~~

示例 2：
~~~
输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1
~~~

## 题目大意

给定一个升序数组，并给定一个target值，求解target对应的数组下标

## 解题思路

**方法一：**
1. 二分查找，对半查找
2. 使用双指针low，high进行对半查找
