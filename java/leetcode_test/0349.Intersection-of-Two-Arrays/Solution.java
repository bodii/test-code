package intersection_of_two_arrays;

import java.util.ArrayList;

class Solution {
    public int[] intersection(int[] nums1, int[] nums2) {
        ArrayList<Integer> arr = new ArrayList<>();

        for (int i = 0; i < nums1.length; i ++) 
            for (int j = 0; j < nums2.length; j ++)
                if (nums2[j] == nums1[i] && !arr.contains(nums1[i]))
                    arr.add(nums1[i]);  

        int[] result = new int[arr.size()];
        for (int i = 0; i < arr.size(); i++) 
            result[i] = arr.get(i);
            
        return result;
    }
}

class Test {
    public static void main(String[] args) {
        int[] nums1 = {1,2,2,1}, nums2 = {2,2};
        int[] result1 = (new Solution()).intersection(nums1, nums2);
        for (int r1 : result1)
            System.out.println(r1);

        int[] nums3 = {4,9,5}, nums4 = {9,4,9,8,4};
        int[] result2 = (new Solution()).intersection(nums3, nums4);
        for (int r2 : result2)
            System.out.println(r2);
    }
}