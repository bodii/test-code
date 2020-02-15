package remove_element;

class Solution2 {
    public int removeElement(int[] nums, int val) {
        int len = nums.length;
        for (int i = 0; i < len; i ++){
            if (nums[i] != val)
                continue; // 如果当前索引位上的值与比较值不相等则跳过

            // 查询最近的不相等的索引
            int r = findNonEqualsIndex(i, len-1, nums, val);
            if (-1 == r)  // 如果查询不到，则返回开始到当前索引位的长度
                return len - (len - i);
            swap(i, r, nums); // 否则交换两个索引位的值
        }
            
        return len;
    }

    /**
     * 查找当前索引后与比较数不相等的最近的索引
     * 
     * @param l 开始索引
     * @param r 结束索引
     * @param nums 原数组
     * @param val 比较的值
     * @return 不相等的最近的索引
     */
    private int findNonEqualsIndex(int l, int r, int[] nums, int val) {
        while(l < r) 
            if (nums[++l] != val) 
                return l;

        return -1;
    }

    /**
     * 对两个索引上的元素进行位置交换
     * 
     * @param l 左边的索引
     * @param r 右边的索引
     * @param nums 原数组
     */
    private void swap(int l, int r, int[] nums) {
        int temp = nums[r];
        nums[r] = nums[l];
        nums[l] = temp;
    }
}

class Test2 {
    public static void main(String[] args) {
        Solution2 solution = new Solution2();

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

        int[] nums3 = {3, 3};
        int val3 = 3;
        int len3 = solution.removeElement(nums3, val3);
        System.out.println("len: " + len3);
        for (int i = 0; i < len3; i ++)
            System.out.print(nums3[i] + ", ");

        System.out.println();

    }
}