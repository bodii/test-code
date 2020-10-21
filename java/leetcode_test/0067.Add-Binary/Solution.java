package add_binary;

class Solution {
    public String addBinary(String a, String b) {
        int alen = a.length();
        int blen = b.length();
        int len = alen > blen ? alen : blen;
        int tmp = 0;
        StringBuffer result = new StringBuffer();
        for (int i = 0; i < len; i ++) {
            if (i < alen) {
                tmp += a.charAt(alen - i - 1) - '0';
            }

            if (i < blen) {
                tmp += b.charAt(blen - i - 1) - '0';
            }

            result.insert(0, tmp % 2);
            tmp /= 2;
        }

        if (tmp > 0) {
            result.insert(0, 1);
        }

        return result.toString();
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        String a = "0";
        String b = "0";
        System.out.println(s.addBinary(a, b));
    }
}
