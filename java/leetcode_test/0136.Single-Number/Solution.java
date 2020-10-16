package single_number;

class Solution {
    public int singleNumber(int[] nums) {
        /*
        int result = nums[0];
        for (int i = 1; i < nums.length; i ++) {
            System.out.printf("result:%d num:%d xor:%d\n", result, nums[i], result ^ nums[i]);
            result ^= nums[i];
        }
        return result;
        */
        for (int i = 1; i < nums.length; i ++) {
            nums[0] ^= nums[i];
        }
        return nums[0];
        
    }

    public static void main(String[] args) {
        int[] nums = {4, 1, 2, 1, 2};
        Solution s = new Solution();
        System.out.println(s.singleNumber(nums));
    }
}