package range_sum_query_mutable;

import range_sum_query_mutable.Merger;

/**
 * SegmentTree
 * 线段树（区间树）类
 * 关心树的线段或区间
 */
public class SegmentTree<E> {
    private E[] data; // 原始数组内容
    private E[] tree; // 线段树数组内容
    private Merger<E> merger; // 融合方法(Merger接口类的实现类merger方法)

    /**
     * 构造函数
     * 用于初始化线段树和原始数组
     * @param arr 原数组
     * @param merger Merger接口类的实现类merger方法
     */
    public SegmentTree(E[] arr, Merger<E> merger) {
        this.merger = merger;

        data = (E[])new Object[arr.length];

        for (int i = 0; i < arr.length; i++ ) 
            data[i] = arr[i];

        // 线段树大约用到的空间是4 * 原数组的长度;
        tree = (E[])new Object[4 * arr.length]; 
        bindSegmentTree(0, 0, data.length -1); // 
    }

    /**
     * 递归绑定线段树内容
     * 在treeIndex的位置创建区间
     * 
     * @param treeIndex 当前索引
     * @param l 左孩子索引
     * @param r 右孩子索引
     */
    private void bindSegmentTree(int treeIndex, int l, int r) {
        if (l == r) {
            tree[treeIndex] = data[l]; // 最后一个元素
            return;
        }

        int leftTreeIndex = leftChild(treeIndex); // 左孩子索引
        int rightTreeIndex = rightChild(treeIndex); // 右孩子索引
        int mid = l + (r - l) / 2; // 中间索引
        
        // 创建左子树(区间: 左边到中间的距离)
        bindSegmentTree(leftTreeIndex, l, mid); 
        // 创建右子树(区间: 中间+1到右边的距离)
        bindSegmentTree(rightTreeIndex, mid + 1, r); 

        // 当前树索引 ＝ 将左孩子树的融合值与右孩子的树的融合值 融合
        tree[treeIndex] = merger.merge(tree[leftTreeIndex], tree[rightTreeIndex]);
    }

    /**
     * 查询两个索引区间的结果
     * 
     * @param queryL 左开始
     * @param queryR 右结束
     * @return 区间结果
     */
    public E query(int queryL, int queryR) {
        ensureIndex(queryL);
        ensureIndex(queryR);
        if (queryL > queryR)
            throw new IllegalArgumentException("Index is illegal.");

        return query(0, 0, data.length - 1, queryL, queryR);
    }

    /**
     * 递归查询两个区间索引在线段树中的索引位置，并返回区间的值
     * 
     * @param treeIndex 当前索引
     * @param l 左开始区间
     * @param r 右开始区间
     * @param queryL 要查询的左开始
     * @param queryR 要查询的右结束
     * @return 区间的结果
     */
    private E query(int treeIndex, int l, int r, int queryL, int queryR) {
        // 如果左开始区间等于要查询的左开始，并且右开始区间等于要查询的右开始
        // 返回当前左区间到右区间，也就是当前值
        if (l == queryL && r == queryR)
            return tree[treeIndex]; 

        int mid = l + (r - l) / 2; // 中间区间
        int leftTreeIndex = leftChild(treeIndex);
        int rightTreeIndex = rightChild(treeIndex);

        // 如果要查询的左开始大于中间区间加1
        // 则从右区间的开始查询，反之则从左区间查询
        if (queryL >= mid + 1)
            return query(rightTreeIndex, mid + 1, r, queryL, queryR);
        else if (queryR <= mid)
            return query(leftTreeIndex, l, mid, queryL, queryR);

        // 如果要查询的左开始在左区间的右边
        // 要查询的右开始在右区间的左边
        E leftResult = query(leftTreeIndex, l, mid, queryL, mid);
        E rightResult = query(rightTreeIndex, mid + 1, r,  mid + 1, queryR);

        return merger.merge(leftResult, rightResult);
    }

    /**
     * 更新线段树某个索引的值
     * 
     * @param index 指定的索引
     * @param e 新元素值
     */
    public void set(int index, E e) {
        ensureIndex(index);

        data[index] = e; // 更新数组这个索引的值
        // 更新线段树的值
        set(0, 0, data.length - 1, index, e);
    }

    /**
     * 递归更新以treeIndex为根的线段树中更新index的值为e
     * 
     * @param treeIndex 当前索引
     * @param l 左开始
     * @param r 右结束
     * @param index 指定要更新的索引
     * @param e 新元素的值
     */
    private void set(int treeIndex, int l, int r, int index, E e) {
        if (l == r) {
            tree[treeIndex] = e;
            return;
        }

        int mid = l + (r -l) / 2;
        int leftTreeIndex = leftChild(treeIndex);
        int rightTreeIndex = rightChild(treeIndex);

        if (index >= mid + 1)
            set(rightTreeIndex, mid + 1, r, index, e);
        else
            set(leftTreeIndex, l, mid, index, e);

        tree[treeIndex] = merger.merge(tree[leftTreeIndex], tree[rightTreeIndex]);
    }

    /**
     * 获取原始数组的元素个数
     * 
     * @return 元素个数
     */
    public int getSize() {
        return data.length;
    }

    /**
     * 检查原始数组是否为空
     * 
     * @return 是否为空
     */
    public boolean isEmpty() {
        return data.length == 0;
    }

    /**
     * 获取原始数组的指定索引的元素
     * 
     * @param index 索引
     * @return 元素
     */
    public E get(int index) {
        ensureIndex(index);

        return data[index];
    }

    /**
     * 获取线段树指定节点的左孩子索引
     * 
     * @param index 当前节点的索引
     * @return 左孩子索引
     */
    private int leftChild(int index) {
        return index * 2 + 1;
    }

    /**
     * 获取线段树指定节点的右孩子索引
     * 
     * @param index 当前节点的索引
     * @return 右孩子索引
     */
    private int rightChild(int index) {
        return index * 2 + 2;
    }

    /**
     * 检测原生数组索引的范围是否合法
     * 
     * @param index 索引
     * @throws IllegalArgumentException
     */
    private void ensureIndex(int index) {
        if (index < 0 || index >= data.length)
            throw new IllegalArgumentException("Failed,index is out of range!");
    }

    /**
     * 将当前线段树转成字符串
     * 
     * @return 转换成的字符串内容
     */
    @Override
    public String toString() {
        StringBuilder result = new StringBuilder();
        result.append("[");

        for (int i = 0; i < tree.length; i ++) {
            if (tree[i] != null)
                result.append(tree[i]);
            else    
                result.append("Null");

            if (i != tree.length - 1)
                result.append(", ");
        }

        result.append("]");

        return result.toString();
    }
}