package two_sum;

import java.util.*;

class Solution {
    public int[] twoSum(int[] nums, int target) {
        int[] result;
        int i = 0;
        int end;
        for (int j = i + 1; j < nums.length; j++) {
            if (nums[i] + nums[j] != target) {
                
            }
        } 
    }
}

/**
 * 
 * @param args
 */
public class TestTwoSum {
    public static void main(String[] args) {
        Solution s = new Solution();
        int[] nums = new int[]{2, 7, 11, 15};
        // int[] nums = new int[]{};
        int target = 26;

        int[] result = s.twoSum(nums, target);
        for (int r : result)
            System.out.print(r + ", ");
        System.out.println();
        
    }
}