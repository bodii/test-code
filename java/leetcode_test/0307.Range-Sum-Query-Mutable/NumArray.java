package range_sum_query_mutable;

import range_sum_query_mutable.SegmentTree;

class NumArray {

    private SegmentTree<Integer> segment;

    public NumArray(int[] nums) {
        Integer[] data = new Integer[nums.length];
        for (int i = 0; i < nums.length; i ++)
            data[i] = nums[i];

        segment = new SegmentTree<>(data, (a, b) -> a + b);
    }
    
    public void update(int i, int val) {
        segment.set(i, val);
    }
    
    public int sumRange(int i, int j) {
        return segment.query(i, j);
    }
}

/**
 * Your NumArray object will be instantiated and called as such:
 * NumArray obj = new NumArray(nums);
 * obj.update(i,val);
 * int param_2 = obj.sumRange(i,j);
 */

 class Test {
     public static void main(String[] args) {
        int[] nums = {1, 3, 5};
        NumArray obj = new NumArray(nums);
        System.out.println("before: " + obj.sumRange(0, 2));
        obj.update(1, 2);
        System.out.println("updated: " + obj.sumRange(0, 2));
        
     }
 }