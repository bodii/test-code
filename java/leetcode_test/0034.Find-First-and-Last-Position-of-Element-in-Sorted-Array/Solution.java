package find_first_and_last_position_of_element_in_sorted_array;

class Solution {
    public int[] searchRange(int[] nums, int target) {

        if (nums.length == 1 && target == nums[0])
            return new int[]{0, 0};

        int k = 0;
       for (int i = 0; i < nums.length; i++)
            if (nums[i] == target)
                nums[k++] = i;

        if (k == 0)
            return new int[]{-1, -1};

        return Arrays.copyOfRange(nums, 0, k);
    }
}


class Test {
    public static void main(String[] args) {
        Solution solution = new Solution();

        int[] nums1 = {5,7,7,8,8,10};
        int target1 = 8;
        int[] result1 = solution.searchRange(nums1, target1);
        for (int r : result1) 
            System.out.println(r);

        System.out.println();
        int[] nums2 = {5,7,7,8,8,10};
        int target2 = 6;
        int[] result2 = solution.searchRange(nums2, target2);
        for (int r : result2) 
            System.out.println(r);
        
    }
}