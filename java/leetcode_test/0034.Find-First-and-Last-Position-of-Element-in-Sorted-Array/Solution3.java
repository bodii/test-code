package find_first_and_last_position_of_element_in_sorted_array;

class Solution3 {
    public int[] searchRange(int[] nums, int target) {

        int[] result = { -1, -1 };
        search(0, nums.length-1, nums, target, result);
        return result;
    }

    private void search(int l, int r , int[] nums, int target, int[] result) {
        if (l > r) 
            return;

        int mid =  l + (r - l) / 2;
        if (target == nums[mid]) {
            if (result[0] > mid || result[0] == -1)
                result[0] = mid;

            if (result[1] < mid)
                result[1] = mid;
        }

        if (target >= nums[l] && target <= nums[mid])
            search(l, mid - 1, nums, target, result);
        
        if (target >= nums[mid] && target <= nums[r])
            search(mid + 1, r, nums, target, result);
    }
}

class Test3 {
    public static void main(String[] args) {
        Solution3 solution = new Solution3();

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