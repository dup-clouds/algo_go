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

1. 给定一个完美二叉树
2. 设置二叉树右侧节点的指针next

## 解题思路

**方法一：**

1. 使用迭代法
2. 按层遍历二叉树
3. 遍历时设置对应的next指针
4. 需要每次更新当前节点进行设置，当前节点即迭代到的节点指针

**方法二：**

1. 递归遍历法
2. 两种形态：若root.left节点不为空，则将left的next指向right；紧跟着root.next不为空，则设置right->next=next.left
3. 递归左节点
4. 递归右节点