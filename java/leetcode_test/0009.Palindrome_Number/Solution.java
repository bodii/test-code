package plaindrome_number;

class Solution {
    public boolean isPalindrome(int x) {
        if (x < 0 ) return false;
        if (x < 10) return true;

        String tmp = Integer.toString(x);
        int len = tmp.length();
        int index = len % 2 == 0 ? len / 2 : len / 2 + 1;

        for (int i = 0; i < index; i++) {
            if (tmp.charAt(i) != tmp.charAt(len - 1 - i)) {
                return false;
            }
        }
        return true;

    }
}


class Test {
    public static void main(String[] args) {
        Solution s = new Solution();
        int x = 67176;
        System.out.println( x + " : 是否是回文数: " + s.isPalindrome(x));
    }
}