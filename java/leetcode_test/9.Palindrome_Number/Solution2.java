package plaindrome_number;

class Solution2 {
    public boolean isPalindrome(int x) {
        if (x < 0 ) return false;
        if (x < 10) return true;

        int len = Integer.toString(x).length();
        int index = len % 2 == 0  ? len / 2 : len / 2 + 1;
        int v = 0;
        while(index > 0) {
            int left = (x / ((int)Math.pow(10, len-v -1))) % 10;
            int right = (x / (int)Math.pow(10, v)) % 10;
            if ( left != right) {
                return false;
            }

            index --;
            v ++;
        }

        return true;
    }
}


class Test2 {
    public static void main(String[] args) {
        Solution2 s = new Solution2();
        int x = 12321;
        System.out.println( x + " : 是否是回文数: " + s.isPalindrome(x));
    }
}