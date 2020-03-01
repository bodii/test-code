package union_find;

import union_find.UnionFindable;

/**
 * UnionFind class
 * 并查集类，又称作Quick Find
 * 时间复杂度O(1)
 */
public class UnionFind  implements UnionFindable {
    private int[] id; // 编号的数组集合

    /**
     * 初始化并查集实例
     * @param size 要存储的元素个数
     */
    public UnionFind(int size) {
        id = new int[size];

        // 初始化并查集内容
        for (int i = 0; i < size; i ++)
            id[i] = i;
    }

    @Override
    public int getSize() {
        return id.length;
    }

    /**
     * 查询元素p所对应的集合编号
     * 
     * @param p 元素
     * @return 集合编号
     */
    private int find(int p) {
        if (p < 0 && p >= id.length)
            throw new IllegalArgumentException("p is out of bound.");

        return id[p];
    }

    /**
     * 检查两个元素是否在一个集合编号内
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
     * 合并元素p和元素q所属的集合
     * 
     * @param p 元素
     * @param q 元素
     */
    @Override
    public void unionElements(int p, int q) {
        int pId = find(p);
        int qId = find(q);

        if (pId == qId) 
            return ;

        // 将所有pId编号设置成qId
        for (int i = 0 ; i < id.length; i ++)
            if (id[i] == pId)
                id[i] = qId;
    }
}