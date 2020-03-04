package union_find;

import union_find.UnionFindable;

/**
 * UnionFind3 class
 * 并查集类
 * Quick union 快速合并
 * 维护一个size 
 * 时间复杂度O(h) h是元素在该树的深度
 */
public class UnionFind3  implements UnionFindable {
    private int[] parent; // 编号的数组集合父节点
    private int[] sizes; // 每个节点对应元素的个数

    /**
     * 初始化并查集实例
     * @param size 要存储的元素个数
     */
    public UnionFind3(int size) {
        parent = new int[size];
        sizes = new int[size];

        // 初始化并查集内容
        for (int i = 0; i < size; i ++) {
            parent[i] = i;
            sizes[i] = 1;
        }
    }

    @Override
    public int getSize() {
        return parent.length;
    }

    /**
     * 查询元素p所对应的根节点编号
     * 
     * @param p 元素
     * @return 集合编号
     */
    private int find(int p) {
        if (p < 0 && p >= parent.length)
            throw new IllegalArgumentException("p is out of bound.");

        while (p != parent[p])
            p = parent[p];

        return p;
    }

    /**
     * 检查两个元素是否两个根节点编号相同
     * 
     * @param p 元素
     * @param q 元素
     * @return 是否在一个集合编号内
     */
    @Override
    public boolean isConnected(int p, int q) {
        return find(p) == find(q);
    }

    /**
     * 合并元素p和元素q所属的根节点
     * 
     * @param p 元素
     * @param q 元素
     */
    @Override
    public void unionElements(int p, int q) {
        int pRoot = find(p);
        int qRoot = find(q);

        if (pRoot == qRoot) 
            return ;

        // 判断，如果pRoot上节点个数小于qRoot,
        if (sizes[pRoot] < sizes[qRoot]) {
            // 将qRoot节点转移到pRoot节点上
            parent[pRoot] = qRoot;
            // 同时，将qRoot节点个数加上pRoot节点个数
            sizes[qRoot] += sizes[pRoot];
        } else {
            parent[qRoot] = pRoot;
            sizes[pRoot] += sizes[qRoot];
        }
    }
}