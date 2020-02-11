package intersection_of_two_arrays;

import java.util.ArrayList;

class Solution4 {
    public int[] intersection(int[] nums1, int[] nums2) {
        int maxLen = nums1.length > nums2.length ? nums1.length : nums2.length;
        int minLen = nums1.length < nums2.length ? nums1.length : nums2.length;
        ArrayList<Integer> list = new ArrayList<>(maxLen);
        ArrayList<Integer> list2 = new ArrayList<>(minLen);

        for (int i = 0; i < nums1.length; i ++) 
            if (!list.contains(nums1[i]))
                list.add(nums1[i]);

        for (int i = 0; i < nums2.length; i ++)
            if (list.contains(nums2[i]) && !list2.contains(nums2[i]))
                list2.add(nums2[i]);

        int[] result = new int[list2.size()];
        for (int i = 0; i < list2.size(); i ++) 
            result[i] = list2.get(i);

        return result;
    }
}

class Test4 {
    public static void main(String[] args) {
        int[] nums1 = {1,2,2,1}, nums2 = {2,2};
        int[] result1 = (new Solution4()).intersection(nums1, nums2);
        for (int r1 : result1)
            System.out.println(r1);

        int[] nums3 = {4,9,5}, nums4 = {9,4,9,8,4};
        int[] result2 = (new Solution4()).intersection(nums3, nums4);
        for (int r2 : result2)
            System.out.println(r2);
    }
}