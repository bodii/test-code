package binary_heap;

import java.util.Random;

import binary_heap.MaxHeap;

public class TestMaxHeap2 {
    public static void main(String[] args) {

        int n = 1000000;
        Random random = new Random();
        Integer[] testData = new Integer[n];
        for (int i = 0; i < n ;i ++)
            testData[i] = random.nextInt(Integer.MAX_VALUE);
        
        double time1 = testHeap(testData, false);
        System.out.println("Without heapify: " + time1 + "s");

        double time2 = testHeap(testData, true);
        System.out.println("With heapify: " + time2 + "s");
    }

    /**
     * 用于测试最大二叉堆用时的函数
     * 
     * @param testData 一个数组
     * @param isHeapIfy 是否是堆类型
     * @return 用时，单位秒
     */
    private static double testHeap(Integer[] testData, boolean isHeapIfy) {
        long startTime = System.nanoTime();

        MaxHeap<Integer> heap;
        if (isHeapIfy) {
            heap = new MaxHeap<>(testData);
        } else {
            heap = new MaxHeap<>();
            for (int data : testData)
                heap.add(data);      
        }

        // 声明一个数组，将最大二叉堆中每次取出最大值
        int[] arr = new int[testData.length];
        for (int i = 0; i < testData.length; i ++)
            arr[i] = heap.extractMax();

        // 检测前一个元素是否大于后一个元素，如果不大于，则抛出异常
        for (int i = 1; i < arr.length; i ++)
            if (arr[i - 1] < arr[i])
                throw new IllegalArgumentException("error.");

        System.out.println("Test MaxHeap completed.");

        long endTime = System.nanoTime();

        return (endTime - startTime) / 1000000000.0; // 返回单位秒

    }
}