package segment_tree;

import segment_tree.Merger;

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