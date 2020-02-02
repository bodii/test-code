package priority_queue;

import priority_queue.Array;

/**
 * 最大二叉堆类
 * 
 * @param <E>
 */
public class MaxHeap<E extends Comparable<E>> {
    private Array<E> data;

    /**
     * 构造函数，使用默认容量
     */
    public MaxHeap() {
        data = new Array<>();
    }

    /**
     * 构造函数
     * 将用户传入的数组最大二叉堆化(heapify)
     * 
     * @param arr 传入的数组
     */
    public MaxHeap(E[] arr) {
        data = new Array<>(arr);  // 实例化这个数组为一个堆

        // 遍历这个堆，从最后一个元素索引的父级索引开始
        for (int i = parent(arr.length -1); i >=0; i --)
            siftDown(i); // 逐一做是父亲元素大于孩子元素的比较交换(数据下沉)
    }

    /**
     * 获取最大二叉堆的元素个数
     * 
     * @return 元素个数
     */
    public int size() {
        return data.getSize();
    }

    /**
     * 检测最大二叉堆是否为空
     * 
     * @return 是否为空
     */
    public boolean isEmpty() {
        return data.isEmpty();
    }

    /**
     * 获取当前索引的父级索引
     * 
     * @param index 当前索引
     * @return 父级索引
     */
    private int parent(int index) {
        isOutRange(index);

        return (index -1) / 2;
    }

    /**
     * 获取当前索引的左孩子索引
     * 
     * @param index 当前索引
     * @return 左孩子索引
     */
    private int leftChild(int index) {
        return index * 2 +  1;
    }

    /**
     * 获取当前索引的右孩子索引
     * 
     * @param index 当前索引
     * @return 右孩子索引
     */
    private int rightChild(int index) {
        return index * 2 + 2;
    }

    /**
     * 检测当前索引是否越界
     * 
     * @param index 当前索引
     * @throws IllegalArgumentException
     */
    private void isOutRange(int index) {
        if (index < 0 || index > data.getSize() - 1)
            throw new IllegalArgumentException("index is out range!");
    }

    /**
     * 添加元素
     * 如果当前元素大于父级元素，则将当前元素与上级元素交换
     * 直到上级元素大于新添加的元素为止
     * 
     * @param e 要添加的元素
     */
    public void add(E e) {
        data.addLast(e);
        siftUp(data.getSize() - 1); // 如果是最小二叉堆，则使用siftDown
    }

    /**
     * 将当前索引的元素满足最大二叉堆属性
     * 数据的上升
     * @param index 当前元素的索引
     */
    private void siftUp(int index) {
        while (index > 0 && data.get(parent(index)).compareTo(data.get(index)) < 0) {
            data.swap(index, parent(index));
            index = parent(index);
        }
    }

    /**
     * 取出最大二叉堆中最大的值
     * 
     * @return 最大的值
     */
    public E extractMax() {
        E max = findMax();

        data.swap(0, data.getSize() - 1); // 将头元素放在末尾
        data.removeLast(); // 删除最后一个元素
        siftDown(0); //将头元素做下沉 // 如果是最小二叉堆，则使用siftUp

        return max;
    }

    /**
     * 查找最大二叉堆中的最大元素
     * 
     * @return 最大的元素，也就是根元素
     */
    public E findMax() {
        if (data.getSize() == 0)
            throw new IllegalArgumentException("index is out range! elements is empty!");

        return data.get(0);
    }

    /**
     * 将当前索引的元素满足最大二叉堆属性
     * 数据的下沉
     * 
     * @param index 当前索引
     */
    private void siftDown(int index) {
        while (leftChild(index) < data.getSize()) {
            int j = leftChild(index); // index 索引的左孩子
            // 如果左孩子＋1，也就是右孩子比总元素数小，并且右孩子比左孩子的元素大
            if (j + 1 < data.getSize() && data.get(j + 1).compareTo(data.get(j)) > 0) 
                j = rightChild(index); // 则将索引变成右孩子

            // 如果当前索引的值大于孩子索引，说明已经满足最大二叉堆，则结束循环
            if (data.get(index).compareTo(data.get(j)) >= 0)
                break; 

            // 此时，data[j],是左右孩子中最大的值
            // 此时，将传入的索引与这个最大值索引进行交换
            data.swap(index, j);
            index = j; // 进行下一轮循环
        }
    }

    /**
     * 取最大二叉堆中最大的元素，并用新元素替换
     * 且要满足最大二叉堆的属性
     * 
     * @param e 新元素
     * @return 最大元素
     */
    public E replace(E e) {
        E max = findMax();

        data.set(0, e);
        siftDown(0);

        return max;
    }

    /**
     * 将当前最大二叉堆转换成字符串
     * 
     * @return 返回转换后的字符串
     */
    @Override public String toString() {
        StringBuilder result = new StringBuilder();
        result.append(String.format("Max heap : size = %d%n", data.getSize()));
        result.append("[");
        for (int i = 0; i < data.getSize(); i ++)
            result.append(i + " : " + data.get(i) + ", ");

        result.replace(result.length() - 2, result.length(), "]");
        
        return result.toString();
    }
    
} 