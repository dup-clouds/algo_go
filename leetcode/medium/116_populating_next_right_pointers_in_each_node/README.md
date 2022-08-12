# [116. populating-next-right-pointers-in-each-node](https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/)

## 题目

**填充每个节点的下一个右侧节点指针**

给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

~~~
struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
~~~

填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
初始状态下，所有next 指针都被设置为 NULL

示例 1：
<img src="https://assets.leetcode.com/uploads/2019/02/14/116_sample.png">

~~~
输入：root = [1,2,3,4,5,6,7]
输出：[1,#,2,3,#,4,5,6,7,#]
解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。序列化的输出按层序遍历排列，同一层节点由 next 指针连接，'#' 标志着每一层的结束。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/populating-next-right-pointers-in-each-node
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
~~~

示例 1：

~~~
输入：root = []
输出：[]
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