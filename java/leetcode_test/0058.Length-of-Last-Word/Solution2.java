package length_of_last_word;

class Solution2 {
    public int lengthOfLastWord(String s) {
        int last = s.lastIndexOf(' ');
        if (last == -1)
            return s.length();
        return s.length() -1 - last;
    }
}

class Test2 {
    public static void main(String[] args) {
        Solution2 solution = new Solution2();
        String s = "Hello World";
        System.out.println(solution.lengthOfLastWord(s));

        String s1 = "a";
        System.out.println(solution.lengthOfLastWord(s1));

        String s2 = " ";
        System.out.println(solution.lengthOfLastWord(s2));
    }
}