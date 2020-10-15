package climbing;

class Solution2 {
    public int climbStairs(int n) {
        if (n <= 2) {
            return n;
        }

        int first = 1, second = 2, tmp = 0;
        while (n-- > 2) {
            tmp = first + second;
            first = second;
            second = tmp;
        }
        return tmp;
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        System.out.println(s.climbStairs(20));
    }
}
