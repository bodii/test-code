package length_of_last_word;

class Solution3 {
    public int lengthOfLastWord(String s) {
        int right = s.length();
        int result = 0;
        while (--right >= 0) {
            if (result == 0 && ' ' == s.charAt(right))
                continue;
            if (' ' == s.charAt(right))
                break;
            else
                result ++;
        }

        return result;
    }
}

class Test3 {
    public static void main(String[] args) {
        Solution3 solution = new Solution3();
        String s = "Hello World";
        System.out.println(solution.lengthOfLastWord(s));

        String s1 = "a";
        System.out.println(solution.lengthOfLastWord(s1));

        String s2 = " ";
        System.out.println(solution.lengthOfLastWord(s2));
    }
}