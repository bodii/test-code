package roman_to_integer;

class Solution {
    public int romanToInt(String s) {
        if (s.length() == 0)
            throw new IllegalArgumentException("The string is not empty!");

        int result = 0;
        for (int i = 0; i < s.length() -1; i ++) {
            if (getRomanToInt(s.charAt(i)) >= getRomanToInt(s.charAt(i+1)))
                result += getRomanToInt(s.charAt(i));
            else
                result -=  getRomanToInt(s.charAt(i));
        }
        result += getRomanToInt(s.charAt(s.length()-1));

        return result;
    }

    private int getRomanToInt(char c) {
        switch (c) {
        case 'I':
            return 1;
        case 'V':
            return 5;
        case 'X':
            return 10;
        case 'L':
            return 50;
        case 'C':
            return 100;
        case 'D':
            return 500;
        case 'M':
            return 1000;
        default:
            return 0;
        }
    }
}

class Test {
    public static void main(String[] args) {
        String s1 = "III";
        Solution solution = new Solution();
        System.out.println(solution.romanToInt(s1));

        String s2 = "IV";
        System.out.println(solution.romanToInt(s2));

        String s3 = "IX";
        System.out.println(solution.romanToInt(s3));

        String s4 = "LVIII";
        System.out.println(solution.romanToInt(s4));

        String s5 = "MCMXCIV";
        System.out.println(solution.romanToInt(s5));

    }
}