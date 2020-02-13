package remove_element;

class Solution {
    public int removeElement(int[] nums, int val) {
        int len = nums.length;
        for (int i = 0; i < len; i ++) {
            if (nums[i] == val) {
                moveEle(i, len--, nums);
            }
        }

        return len;
    }

    private void moveEle(int l , int r, int[] nums) {
        if (l == r)
            return ;

        for (int i = l; i < r-1; i ++)
            nums[i] = nums[i+1];
    }
}

class Test {
    public static void main(String[] args) {
        Solution solution = new Solution();

        int[] nums = {3,2,2,3};
        int val = 3;

        int len = solution.removeElement(nums, val);
        System.out.println("len: " + len);
        for (int i = 0; i < len; i ++)
            System.out.print(nums[i] + ", ");

        System.out.println();


        int[] nums2 = {0,1,2,2,3,0,4,2};
        int val2 = 2;

        int len2 = solution.removeElement(nums2, val2);
        System.out.println("len: " + len2);
        for (int i = 0; i < len2; i ++)
            System.out.print(nums2[i] + ", ");

        System.out.println();

    }
}