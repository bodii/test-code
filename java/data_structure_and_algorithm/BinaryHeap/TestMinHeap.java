package binary_heap;

import java.util.Random;

import binary_heap.MinHeap;

/**
 * 最小二叉堆的测试类
 */
public class TestMinHeap {
    public static void main(String[] args) {
        MinHeap<Integer> heap = new MinHeap<>();

        int n = 1000000;
        Random random = new Random();
        for (int i = 0; i < n; i ++) 
            heap.add(random.nextInt(Integer.MAX_VALUE));

        System.out.println(heap);
        System.out.println("size: " + heap.size());

        int[] arr = new int[n];
        for (int i = 0; i < n; i ++) {
            arr[i] = heap.extractMin(); // 将最大的元素取出
        }

        for (int i = 1; i < n; i ++) 
            if (arr[i - 1] > arr[i]) // 如果前一个元素小于后一个元素，则说明每次取出的最大值有误
                throw new IllegalArgumentException("Error.");

        System.out.println("Test MinHeap completed.");

    }
}