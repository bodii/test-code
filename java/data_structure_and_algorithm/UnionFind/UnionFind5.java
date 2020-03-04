package union_find;

import union_find.UnionFindable;

/**
 * 并查集类
 * 路径压缩 path  compression
 */
public class UnionFind5 implements UnionFindable {
    int[] parent; // 父节点集合
    int[] rank; // 节点的层数

    public UnionFind5(int size) {
        parent = new int[size];
        rank = new int[size];

        for (int i = 0; i < size; i ++) {
            parent[i] = i;
            rank[i] = 1;
        }
    }   

    @Override
    public int getSize() {
        return parent.length;
    }

    /**
     * 查询元素的根节点
     * 
     * @param p 元素
     * @return 根节点
     */
    private int find(int p) {
        if (p < 0 || p >= parent.length)
            throw new IllegalArgumentException("p is out of bound.");
        try {
            while(p != parent[p] && p < parent.length) {
                // 父级 ＝ 父级的父级
                parent[p] = parent[parent[p]];
                p = parent[p];
            }
        } catch (Exception e) {
            //TODO: handle exception
        }

        return p;
    }

    /**
     * 检测两个元素的根节点是否相同
     * 或者说 检测两个元素是否连接 
     * 
     * @param p 元素
     * @param q 元素
     * @return 是否是同一个编号的根节点
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

        if (rank[pRoot] < rank[qRoot])
            parent[pRoot] = qRoot;
        else if (rank[pRoot] > rank[pRoot])
            parent[qRoot] = pRoot;
        else {
            parent[qRoot] = pRoot;
            parent[pRoot] += 1;
        }
    }
}