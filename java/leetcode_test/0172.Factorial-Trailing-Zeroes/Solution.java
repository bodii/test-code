package factorial_trailing_zeroes;

class Solution {
    public int trailingZeroes(int n) {
        int zeroes = 0;

        while (n > 0) {
            n /= 5;
            zeroes += n;
        }

        return zeroes;
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        System.out.println(s.trailingZeroes(100));
    }
}