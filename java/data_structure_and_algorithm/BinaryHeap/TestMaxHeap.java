package binary_heap;

import java.util.Random;

import binary_heap.MaxHeap;

/**
 * 最大二叉堆的测试类
 */
public class TestMaxHeap {
    public static void main(String[] args) {
        MaxHeap<Integer> heap = new MaxHeap<>();

        int n = 1000000;
        Random random = new Random();
        for (int i = 0; i < n; i ++) 
            heap.add(random.nextInt(Integer.MAX_VALUE));

        System.out.println(heap);

        int[] arr = new int[n];
        for (int i = 0; i < n; i ++) {
            arr[i] = heap.extractMax(); // 将最大的元素取出
            // System.out.println(arr[i]);
        }

        for (int i = 1; i < n; i ++) 
            if (arr[i - 1] < arr[i]) // 如果前一个元素小于后一个元素，则说明每次取出的最大值有误
                throw new IllegalArgumentException("Error.");

        System.out.println("Test MaxHeap completed.");

    }
}