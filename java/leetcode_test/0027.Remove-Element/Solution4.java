package remove_element;

class Solution4 {
    public int removeElement(int[] nums, int val) {
        int res = 0;
        int len = nums.length;
        for (int i = 0; i < len; i ++) {
            if (val == nums[i])
                nums[i--] = nums[--len];
            else
                res ++;
        }

        return res;
    }
}

class Test4 {
    public static void main(String[] args) {
        Solution4 solution = new Solution4();

        int[] nums = {3,2,2,3};
        int val = 3;

        int len = solution.removeElement(nums, val);
        System.out.println("len: " + len);
        for (int i = 0; i < len; i ++)
            System.out.print(nums[i] + ", ");

        System.out.println("\n");


        int[] nums2 = {0,1,2,2,3,0,4,2};
        int val2 = 2;

        int len2 = solution.removeElement(nums2, val2);
        System.out.println("len: " + len2);
        for (int i = 0; i < len2; i ++)
            System.out.print(nums2[i] + ", ");

        System.out.println("\n");

        int[] nums3 = {3, 3};
        int val3 = 3;
        int len3 = solution.removeElement(nums3, val3);
        System.out.println("len: " + len3);
        for (int i = 0; i < len3; i ++)
            System.out.print(nums3[i] + ", ");

        System.out.println();

    }
}