package roman_to_integer;

import java.util.ArrayList;

class Solution {
    public int romanToInt(String s) {
        if (s.length() == 0 ) 
            throw new IllegalArgumentException("The string is not empty!");

        ArrayList<Integer> arr = new ArrayList<>(s.length());

        for (int i = s.length() -1; i >= 0; i --) {
            if ( 'I' == s.charAt(i)) {
                int num = 1, index = --i;
                while('I' == s.charAt(index--)) {
                    i --;
                    num ++;
                }
                if ('V' == s.charAt(i)) {
                    i --;
                    num += 5;
                }
                arr.add(num);
            } else if ('V' == s.charAt(i)) {
                int num = 5;
                if ('I' == s.charAt(--i)) {
                    num --;
                    i --;
                }
                arr.add(num);
            } else if ('X' == s.charAt(i)) {
                int num = 1;
                if ('I' == s.charAt(--i)) {
                    num = 9;
                    i --;
                }
                arr.add(num);
            } else if 
        }
    }
}


class Test {
    public static void main(String[] args) {
        String s1 = "III";
        Solution solution = new Solution();
        System.out.println(
            solution.romanToInt(s1)
        );

        String s2 = "IV";
        System.out.println(
            solution.romanToInt(s2)
        );

        String s3 = "IX";
        System.out.println(
            solution.romanToInt(s3)
        );

        String s4 = "LVIII";
        System.out.println(
            solution.romanToInt(s4)
        );

        String s5 = "MCMXCIV";
        System.out.println(
            solution.romanToInt(s5)
        );

    }
}