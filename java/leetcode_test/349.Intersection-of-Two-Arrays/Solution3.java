package intersection_of_two_arrays;

import java.util.ArrayList;
import java.util.TreeSet;

class Solution3 {
    public int[] intersection(int[] nums1, int[] nums2) {
        TreeSet<Integer> arr = new TreeSet<>();

        for (int num : nums1)
            arr.add(num);

        int minLen = nums1.length < nums2.length ? nums1.length : nums2.length;
        ArrayList<Integer> list = new ArrayList<>(minLen);
        for (int num : nums2) {
            if (arr.contains(num)) {
                list.add(num);
                arr.remove(num);
            }
        }

        int[] result = new int[list.size()];
        for (int i = 0; i < list.size(); i++ ) 
            result[i] = list.get(i);

        return result; 
    }
}

class Test3 {
    public static void main(String[] args) {
        int[] nums1 = {1,2,2,1}, nums2 = {2,2};
        int[] result1 = (new Solution3()).intersection(nums1, nums2);
        for (int r1 : result1)
            System.out.println(r1);

        int[] nums3 = {4,9,5}, nums4 = {9,4,9,8,4};
        int[] result2 = (new Solution3()).intersection(nums3, nums4);
        for (int r2 : result2)
            System.out.println(r2);
    }
}