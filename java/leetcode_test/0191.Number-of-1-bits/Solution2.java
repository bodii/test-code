package number_of_bits;

class Solution2 {
    // // you need to treat n as an unsigned value
    public int hammingWeight(int n) {
        int bits = 0;
        while(n != 0) {
            if ((n & 1) == 1) {
                bits ++;
            } 
            n >>= 1;   
        }
        return bits;
    }

    public static void main(String[] args) {
        Solution2 s = new Solution2();
        int n = 00000000000000000000000010000000;
        System.out.println(s.hammingWeight(n));
    }
}
