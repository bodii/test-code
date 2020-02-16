package reange_sum_query_lmmutable;

class NumArray2 {
    private int[] sum;

    public NumArray2(int[] nums) {
        sum = new int[nums.length+1];
        sum[0] = 0;
        for (int i = 1; i < sum.length; i++)
            sum[i] = sum[i-1] + nums[i-1]; 
    }
    
    public int sumRange(int i, int j) {
       return sum[j+1] - sum[i];
    }
}

/**
 * Your NumArray object will be instantiated and called as such:
 * NumArray obj = new NumArray(nums);
 * int param_1 = obj.sumRange(i,j);
 */

 class Test2 {
     public static void main(String[] args) {
        int[] nums = {-2, 0, 3, -5, 2, -1};
        NumArray2 obj = new NumArray2(nums);

        System.out.println(obj.sumRange(0, 2));
        System.out.println(obj.sumRange(2, 5));
        System.out.println(obj.sumRange(0, 5));
     }
 }