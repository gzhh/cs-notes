# RDMS
- MySQL VS PostgreSQL，谁是世界上最成功的数据库？https://mp.weixin.qq.com/s/d_FixQwDDXbJtv3x15G62A


## 数据结构
1. 二叉搜索树（二叉排序树）
    
    https://en.wikipedia.org/wiki/Binary_search_tree
    
    左 < 父 < 右
    
2. AVL 树（一种自平衡 BST）
    
    https://en.wikipedia.org/wiki/AVL_tree
    
    平衡的 BST（叶子高度差小于等于1），会旋转调整树高度来保证平衡
    
3. R.B 树（红黑树，一种自平衡 BST）
    
    [https://en.wikipedia.org/wiki/Red–black_tree](https://en.wikipedia.org/wiki/Red%E2%80%93black_tree)
    
    与 AVL 比，牺牲了部分平衡性，以换取插入/删除时的少量旋转操作，比 AVL 性能更优
    
    典型应用
    
    - 关联数组
    - 集合
4. B-Tree（多路搜索树）
    
    https://en.wikipedia.org/wiki/B-tree
    
    左 < 父 < 右
    
    1. 平衡的，叶子结点高度差小于等于1
    2. 根结点至少有两个孩子
    3. 非根结点的孩子以升序排列
5. B+Tree（B-Tree变体）
    
    https://en.wikipedia.org/wiki/B%2B_tree
    
    1. 内部结点不存储 data，只有 key
    2. 所有 data 都在叶子结点中
    3. 双指针链接结点，方便遍历
    
    应用：MySQL 文件系统
