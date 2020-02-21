package length_of_last_word;

class Solution {
    public int lengthOfLastWord(String s) {
        int result = 0;
        for (int i = s.length() - 1; i >= 0; i --) {
            if (s.charAt(i) == ' ')
                return result; 
            result ++;
        }

        return result;
    }
}

class Test {
    public static void main(String[] args) {
        Solution solution = new Solution();
        String s = "Hello World";
        System.out.println(solution.lengthOfLastWord(s));

        String s1 = "a";
        System.out.println(solution.lengthOfLastWord(s1));

        String s2 = " ";
        System.out.println(solution.lengthOfLastWord(s2));
    }
}