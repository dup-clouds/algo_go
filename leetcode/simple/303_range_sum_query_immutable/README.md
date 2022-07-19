# [83. range-sum-query-immutable](https://leetcode.cn/problems/range-sum-query-immutable/)

## 题目

给定一个整数数组 nums，处理以下类型的多个查询:
计算索引left和right（包含 left 和 right）之间的 nums 元素的 和 ，其中left <= right 实现 NumArray 类：
NumArray(int[] nums) 使用数组 nums 初始化对象
int sumRange(int i, int j) 返回数组 nums中索引left和right之间的元素的 总和 ，包含left和right两点（也就是nums[left] + nums[left + 1] + ... + nums[right])


示例 1：

~~~
输入：
["NumArray", "sumRange", "sumRange", "sumRange"]
[[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
输出：
[null, 1, -1, -3]
~~~

解释：
NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]); <br>
numArray.sumRange(0, 2); // return 1 ((-2) + 0 + 3)  <br>
numArray.sumRange(2, 5); // return -1 (3 + (-5) + 2 + (-1))  <br>
numArray.sumRange(0, 5); // return -3 ((-2) + 0 + 3 + (-5) + 2 + (-1))  <br>

## 题目大意

构造一个NewArray,支持将数组nums存入，并且通过指定方法sumRange(left, right) 获取区间的和

## 解题思路

**方法一：**
1. NewArray 内部使用[]int
2. 构造函数初始化前缀和存入NewArray，其中sums数组长度为n+1，n为原始数组长度，其第一个位置存储0，往后位置存储前面所有位置和
3. sumRange函数通过取前缀和数组中的值进行相减即可得到目标值，通过right+1处的元素-left位置元素可获取sum结果
4. 初始化的时间复杂度为O(N)，每次获取区间sum的时间复杂度为O(1)
