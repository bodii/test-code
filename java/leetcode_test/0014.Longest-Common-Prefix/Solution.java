package longest_common_prefix;

import java.util.ArrayList;

class Solution {
    public String longestCommonPrefix(String[] strs) {
        if (strs.length < 1)
            return "";

        String headStr = strs[0];
        ArrayList<Character> charsArr = new ArrayList<>(headStr.length());
        for(int i = 0; i < headStr.length(); i++)
            charsArr.add(headStr.charAt(i));
        
        int k = 0, i = 0;
        boolean hasContinue = true;
       while(hasContinue && i < headStr.length()){
            for (int j = 0; j < strs.length; j ++) {
                if (i > strs[j].length() -1 ||  !charsArr.get(i).equals(strs[j].charAt(i))) {
                    hasContinue = false;
                    k --;
                    break;
                }
            }
            k ++;
            i ++;
        }

        return headStr.substring(0, k);
    }
}


class Test {
    public static void main(String[] args) {
        String[] strs = {"flower","flow","flight"};

        System.out.println(
            (new Solution()).longestCommonPrefix(strs)
        );

        String[] strs2 = {"ab", "ac"};

        System.out.println(
            (new Solution()).longestCommonPrefix(strs2)
        );
    }
}