package top_k_frequent_elements;

import java.util.LinkedList;
import java.util.List;
import java.util.TreeMap;
import top_k_frequent_elements.PriorityQueue;

class Solution {
    /**
     * 装载一个元素和它的频率的类
     * 
     * @param <E>
     */
    private class Frequency implements Comparable<Frequency> {
        public int e; // 元素
        public int freq; // 出现的频率

        public Frequency(int e, int freq ) {
            this.e = e;
            this.freq = freq;
        }

        /**
         * 当前元素的频率和另一个计算频率类比较函数
         * 
         * @param another 另一个计算频率类
         * @return 比较大小的结果
         */
        @Override 
        public int compareTo(Frequency another) {
            if (freq < another.freq) // 当前频率小于另一个类元素的频率时
                return 1;
            else if (freq > another.freq)
                return -1;
            return 0;
        }
    }

    public List<Integer> topKFrequent(int[] nums, int k) {
        TreeMap<Integer, Integer> map = new TreeMap<>();
        
        for (int num : nums) {
            if (map.containsKey(num)) 
                map.put(num, map.get(num) + 1);
            else 
                map.put(num, 1);
        }

        PriorityQueue<Frequency> queue = new PriorityQueue<>();
        for (int m : map.keySet()) {
            if (queue.getSize() < k)
                queue.enqueue(new Frequency(m, map.get(m)));
            else if (map.get(m) > queue.getFront().freq) {
                queue.dequeue();
                queue.enqueue(new Frequency(m, map.get(m)));
            }
        }

        LinkedList<Integer> list = new LinkedList<>();
        while(!queue.isEmpty())
            list.add(queue.dequeue().e);

        return list;
    }
}

class Test {
    public static void main(String[] args) {
        int[] nums = new int[]{1,1,1,2,2,3};
        int k = 2;

        Solution solution = new Solution();
        System.out.println(solution.topKFrequent(nums, k));
    }
}