package longest_common_prefix;

import java.util.ArrayList;

class Solution2 {
    public String longestCommonPrefix(String[] strs) {
        if (strs.length < 1)
            return "";

        int k = 0, i = 0;
        boolean isOk = true;
        while(i <  strs[0].length()){
            for (int j = 0; j < strs.length; j ++) {
                if (i > strs[j].length() -1 || strs[0].charAt(i) != strs[j].charAt(i)) {
                    isOk = false;
                    break;
                }
            }

            if (isOk == false)
                break;
            k ++;
            i ++;
        }

        return  strs[0].substring(0, k);
    }
}


class Test2 {
    public static void main(String[] args) {
        String[] strs = {"flower","flow","flight"};

        System.out.println(
            (new Solution2()).longestCommonPrefix(strs)
        );

        String[] strs2 = {"aca","cba"};

        System.out.println(
            (new Solution2()).longestCommonPrefix(strs2)
        );

        String[] strs3 = {"dog","racecar","car"} ;

        System.out.println(
            (new Solution2()).longestCommonPrefix(strs3)
        );
    }
}