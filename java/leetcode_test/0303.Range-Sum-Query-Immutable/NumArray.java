package reange_sum_query_lmmutable;

import reange_sum_query_lmmutable.SegmentTree;

class NumArray {

    private SegmentTree<Integer> segmentTree;

    public NumArray(int[] nums) {
        if (nums.length < 1)
            return;

        Integer[] data = new Integer[nums.length];
        for (int i = 0; i < nums.length; i++) 
            data[i] = nums[i];

        segmentTree = new SegmentTree<>(data, (a, b) -> a + b);
    }
    
    public int sumRange(int i, int j) {
        if (segmentTree.isEmpty())
            throw new IllegalArgumentException("Segment tree is empty!");

        return segmentTree.query(i, j);
    }
}

/**
 * Your NumArray object will be instantiated and called as such:
 * NumArray obj = new NumArray(nums);
 * int param_1 = obj.sumRange(i,j);
 */

 class Test {
     public static void main(String[] args) {
        int[] nums = {-2, 0, 3, -5, 2, -1};
        NumArray obj = new NumArray(nums);

        System.out.println(obj.sumRange(0, 2));
        System.out.println(obj.sumRange(2, 5));
        System.out.println(obj.sumRange(0, 5));
     }
 }