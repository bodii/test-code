package binary_heap;

import binary_heap.Array;

/**
 * 最小二叉堆类
 * 所有父级元素的值小于左右孩子的值
 * 
 * @param <E>
 */
public class MinHeap<E extends Comparable<E>> {
    private Array<E> data; // 最小二叉堆的内部数据集

    /**
     * 构造函数
     * 用于实现最小二叉堆对象
     */
    public MinHeap() {
        data = new Array<>();
    }

    /**
     * 构造函数
     *  用于初始化一个满足最小二叉堆的实例对象
     * 
     * @param arr
     */
    public MinHeap(E[] arr) {
        data = new Array<>(arr);

        for (int i = parent(arr.length - 1); i >= 0; i --)
            siftDown(i);
    }

    /**
     * 获取最小二叉树的元素个数
     * 
     * @return 元素个数
     */
    public int size() {
        return data.getSize();
    }

    /**
     * 添加元素 
     * 满足最小二叉堆属性
     * 
     * @param e 要添加的元素
     */
    public void add(E e) {
        data.addLast(e);

        siftUp(data.getSize() -1);
    }

    /**
     * 查看最小元素的值
     *
     * @return 最小元素的值
     */
    public E findMin() {
        isOutRange(0, 0);

        return data.get(0);
    }

    /**
     * 取出最小二叉堆的最小值
     * 
     * @return 最小的值
     */
    public E extractMin() {
        E min = findMin();

        data.swap(0, data.getSize() - 1); // 将头元素放在末尾
        data.removeLast(); // 删除末尾元素
        siftDown(0); // 将当前头元素做下沉处理

        return min;
    }

    /**
     * 替换头元素为新的元素，并将旧的元素返回
     * 
     * @param e 新元素
     * @return 旧元素
     */
    public E replace(E e) {
        E min = findMin();

        data.set(0, e);
        siftDown(0);

        return min;
    }

    /**
     * 获取当前索引的父级索引
     * 
     * @param index 当前索引
     * @return 父级索引
     */
    private int parent(int index) {
        isOutRange(index, 0);

        return (index - 1) / 2;
    }

    /**
     * 获取当前索引的左孩子索引
     * 
     * @param index 当前索引
     * @return 左孩子索引
     */
    private int leftChild(int index) {
        isOutRange(index, 0);

        return index * 2 + 1;
    }

    /**
     * 获取当前索引的右孩子索引
     * 
     * @param index 当前索引
     * @return 右孩子索引
     */
    private int rightChild(int index) {
        isOutRange(index, 0);

        return index * 2 + 2;
    }

    /**
     * 将当前索引的元素与上级所有元素进行比对，并满足最小二叉堆的属性
     * 数据上升
     * 
     * @param index 当前索引
     */
    private void siftUp(int index) {
        while( index > 0 && data.get(index).compareTo(data.get(parent(index))) < 0) {
            data.swap(index, parent(index));
            index = parent(index);
        }
    }

    /**
     * 将当前索引的元素与下级所有元素进行比对，并满足最小二叉堆的属性
     * 数据的下沉
     * 
     * @param index
     */
    private void siftDown(int index) {
        while (leftChild(index) < data.getSize()) {
            int j = leftChild(index);
            if (j + 1 < data.getSize() && data.get(j + 1).compareTo(data.get(j)) < 0)
                j = j + 1;
            
            if (data.get(index).compareTo(data.get(j)) <= 0)
                break;

            data.swap(index, j);
            index = j;
        }
    }

    /**
     * 检测当前索引有没有越界
     * 
     * @param index 当前索引
     * @throws IllegalArgumentException
     */
    private void isOutRange(int index, int minIndex) {
        if (index != 0 && (index < minIndex || index > data.getSize() -1) ) {
            throw new IllegalArgumentException("index is out range!");
        }
    }

    /**
     * 将当前最小二叉堆转换成字符串串
     * 
     * @return 生成转换后的字符串
     */
    @Override public String toString() {
        StringBuilder result = new StringBuilder();

        result.append(String.format("MinHeap : size = %d%n", data.getSize()));
        result.append("[");
        for (int i = 0; i < data.getSize(); i ++)
            result.append(i + ": " + data.get(i) + ", ");
        result.replace(result.length() -2 , result.length(), "]");

        return result.toString();
    }
}