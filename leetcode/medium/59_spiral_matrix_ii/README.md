# [59. spiral-matrix-ii](https://leetcode.cn/problems/spiral-matrix-ii/)
## 题目

给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。


示例 1：

<img src="https://assets.leetcode.com/uploads/2020/11/13/spiraln.jpg">

~~~
输入：n = 3
输出：[[1,2,3],[8,9,4],[7,6,5]]
~~~

示例 2：

~~~
输入：n = 1
输出：[[1]]
~~~

## 题目大意

1. 给定一个n，生成一个n*n的矩阵，且矩阵满足顺时针螺旋从小到大排列关系
2. 顺时针螺旋矩阵并赋值

## 解题思路

**方法一：**
1. 锁定边界：(0, 0) 左上边界，(n-1, n-1) 右下边界,按层遍历，每遍历一次则对应缩小边界
2. 向右，从left到right
3. 向下，从top+1到bottom
4. 向左，从right-1到left
5. 向上，从bottom-1到top+1