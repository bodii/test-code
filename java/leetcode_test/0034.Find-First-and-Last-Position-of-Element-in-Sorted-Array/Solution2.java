package find_first_and_last_position_of_element_in_sorted_array;

class Solution2 {
    public int[] searchRange(int[] nums, int target) {

        int[] result = { -1, -1 };
        int k = 0;
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] == target) {
                result[0] = i;
                break;
            }
        }

        if (result[0] == -1)
            return result;

        for (int i = nums.length - 1; i >= result[0]; i--) {
            if (nums[i] == target) {
                result[1] = i;
                break;
            }
        }

        return result;
    }
}

class Test2 {
    public static void main(String[] args) {
        Solution2 solution = new Solution2();

        int[] nums1 = { 5, 7, 7, 8, 8, 10 };
        int target1 = 8;
        int[] result1 = solution.searchRange(nums1, target1);
        for (int r : result1)
            System.out.println(r);

        System.out.println();
        int[] nums2 = { 5, 7, 7, 8, 8, 10 };
        int target2 = 6;
        int[] result2 = solution.searchRange(nums2, target2);
        for (int r : result2)
            System.out.println(r);

    }
}