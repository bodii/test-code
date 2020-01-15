package remove_duplicates_from_sorted_array;

class Solution2 {
    public int removeDuplicates(int[] nums) {
        if (nums.length == 0) return 0;

        int i = 0;
        for (int j = 1; j < nums.length; j++ ) {
            if (nums[i] != nums[j]) {
                i ++;
                nums[i] = nums[j];
            }
        }
        return i+1;
    }
}
/*
0,0,1,1,1,2,2,3,3,4
0 = 1
0,1,1,1,2,2,3,3,4,4
1 = 1
0,1,1,2,2,3,3,4


*/

class Test2 {
    public static void main(String[] args) {
        Solution2 s = new Solution2();
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