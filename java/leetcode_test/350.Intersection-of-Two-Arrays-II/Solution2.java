package intersection_of_two_arrays_2;

import java.util.Arrays;

class Solution2 {
    public int[] intersect(int[] nums1, int[] nums2) {
        Arrays.sort(nums1);
        Arrays.sort(nums2);

        int i = 0, j = 0, k = 0;
        while(i < nums1.length && j < nums2.length) {
            if (nums1[i] < nums2[j])
                i ++;
            else if (nums1[i] > nums2[j])
                j ++;
            else {
                nums1[k++] = nums1[i++];
                j ++;
            }
        }

        return Arrays.copyOfRange(nums1, 0, k);
    }
}

class Test2 {
    public static void main(String[] args) {
        int[] nums1 = {1,2,2,1}, nums2 = {2,2};
        int[] result1 = (new Solution2()).intersect(nums1, nums2);
        for (int r1 : result1)
            System.out.println(r1);

        int[] nums3 = {4,9,5}, nums4 = {9,4,9,8,4};
        int[] result2 = (new Solution2()).intersect(nums3, nums4);
        for (int r2 : result2)
            System.out.println(r2);
    }
}