# [304. range-sum-query-2d-immutable](https://leetcode.cn/problems/range-sum-query-2d-immutable/)

## 题目

给定一个二维矩阵 matrix，以下类型的多个请求：
- 计算其子矩形范围内元素的总和，该子矩阵的 左上角 为 (row1,col1) ，右下角 为 (row2,col2) 。
- 
实现 NumMatrix 类：

- NumMatrix(int[][] matrix)给定整数矩阵 matrix 进行初始化
- int sumRegion(int row1, int col1, int row2, int col2)返回 左上角 (row1,col1)、右下角(row2,col2) 所描述的子矩阵的元素 总和 。

<img src="https://pic.leetcode-cn.com/1626332422-wUpUHT-image.png">

示例 1：
~~~
输入:
["NumMatrix","sumRegion","sumRegion","sumRegion"]
[[[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]],[2,1,4,3],[1,1,2,2],[1,2,2,4]]
输出:
[null, 8, 11, 12]

解释:
NumMatrix numMatrix = new NumMatrix([[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]);
numMatrix.sumRegion(2, 1, 4, 3); // return 8 (红色矩形框的元素总和)
numMatrix.sumRegion(1, 1, 2, 2); // return 11 (绿色矩形框的元素总和)
numMatrix.sumRegion(1, 2, 2, 4); // return 12 (蓝色矩形框的元素总和)
~~~

## 题目大意

1. 给定一个二维矩阵
2. 输入矩阵内小矩阵左上顶点+右下顶点，求解包围的矩阵和

## 解题思路

**方法一：**
1. 构建二维矩阵和
2. 通过构建arry[m+1, n+1]矩阵存储前缀和，从第[1,1]处开始，其中[0,0]为0
3. 通过包围范围进行加减操作得出[i,j]矩阵和
4. 初始化的时间复杂度+空间复杂度为O(M*N)，每次获取区间sum的时间复杂度为O(1)