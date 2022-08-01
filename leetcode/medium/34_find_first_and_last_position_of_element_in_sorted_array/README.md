# [34. find-first-and-last-position-of-element-in-sorted-array](https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/)

## 题目
在排序数组中查找元素的第一个和最后一个位置</br>
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回[-1, -1]。
你必须设计并实现时间复杂度为O(log n)的算法解决此问题。

示例 1：

~~~
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
~~~
    
示例 2：

~~~
输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
~~~

示例 3：

~~~
输入：nums = [], target = 0
输出：[-1,-1]
~~~

## 题目大意

1. 给定一个升序数组，目标值
2. 查找第一个位置等于目标值及最后一个位置等于目标值的结果

## 解题思路

**方法一：**
1. 二分查找
2. 处理nums[mid]==target处进行特殊处理
3. 分两次进行查找第一个位置和最后一个位置