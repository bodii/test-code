package intersection_of_two_arrays;

import java.util.TreeSet;

class Solution2 {
    public int[] intersection(int[] nums1, int[] nums2) {
        TreeSet<Integer> arr = new TreeSet<>();

        for (int i = 0; i < nums1.length; i ++) 
            for (int j = 0; j < nums2.length; j ++)
                if (nums2[j] == nums1[i])
                    arr.add(nums1[i]);  

        int[] result = new int[arr.size()];
        int i = 0;
        while(!arr.isEmpty())
            result[i++] = arr.pollLast();

        return result;
    }
}

class Test2 {
    public static void main(String[] args) {
        int[] nums1 = {1,2,2,1}, nums2 = {2,2};
        int[] result1 = (new Solution2()).intersection(nums1, nums2);
        for (int r1 : result1)
            System.out.println(r1);

        int[] nums3 = {4,9,5}, nums4 = {9,4,9,8,4};
        int[] result2 = (new Solution2()).intersection(nums3, nums4);
        for (int r2 : result2)
            System.out.println(r2);
    }
}