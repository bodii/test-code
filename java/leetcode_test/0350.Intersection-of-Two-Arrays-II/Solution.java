package intersection_of_two_arrays_2;

import java.util.ArrayList;
import java.util.TreeMap;

class Solution {
    public int[] intersect(int[] nums1, int[] nums2) {
        TreeMap<Integer, Integer> map = new TreeMap<>();

        for (int num : nums1) {
            if (!map.containsKey(num))
                map.put(num, 1);
            else
                map.put(num, map.get(num) +1);
        }

        ArrayList<Integer> arr = new ArrayList<>(nums1.length); 
        for (int num : nums2) {
            if (map.containsKey(num)) {
                arr.add(num);
                map.put(num, map.get(num) - 1);
                if (0 == map.get(num))
                    map.remove(num);
            }
        }

        int[] result = new int[arr.size()];
        for (int i = 0; i < arr.size(); i++) 
            result[i] = arr.get(i);
            
        return result;
    }
}

class Test {
    public static void main(String[] args) {
        int[] nums1 = {1,2,2,1}, nums2 = {2,2};
        int[] result1 = (new Solution()).intersect(nums1, nums2);
        for (int r1 : result1)
            System.out.println(r1);

        int[] nums3 = {4,9,5}, nums4 = {9,4,9,8,4};
        int[] result2 = (new Solution()).intersect(nums3, nums4);
        for (int r2 : result2)
            System.out.println(r2);
    }
}