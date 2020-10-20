package number_of_bits;

class Solution {
    // you need to treat n as an unsigned value
    public int hammingWeight(int n) {
        int result = 0;
        while (n != 0) {
            n &= (n -1); 
            result ++;
        }

        return result;
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        int n = 00000000000000000000000000001011;
        System.out.println(s.hammingWeight(n));
    }
}