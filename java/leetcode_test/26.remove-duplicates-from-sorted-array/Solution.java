package remove_duplicates_from_sorted_array;

class Solution {
    public int removeDuplicates(int[] nums) {
        int len = nums.length;
        int swapLen = 0;
        for (int i = 1; i < len; i++ ) {
            if (nums[i -1] == nums[i]) {
                swapLen++;
            }
        }

        for (int i = 0; i < swapLen; i++) {
            for (int j = 1; j < len-1;j++) {
                if (nums[j-1] == nums[j]) {
                    nums[j] = nums[j+1];
                }
            }
        }

        return len - swapLen;
    }
}


class Test {
    public static void main(String[] args) {
        Solution s = new Solution();
        int[] nums = new int[]{1,1,2,2};
        int len = s.removeDuplicates(nums);
        System.out.println("len: " + len);
        for (int i = 0; i <len; i++)  {
            if (i != len)
                System.out.print(nums[i] + ", ");
            else 
                System.out.print(nums[i] );
        }
        System.out.println();
    }
}