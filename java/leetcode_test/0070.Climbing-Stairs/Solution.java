package climbing;

class Solution {
    public int climbStairs(int n) {
        if (n <= 2) {
            return n;
        }

        return climbStairs(n - 1) + climbStairs(n - 2);
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        System.out.println(s.climbStairs(10));
    }
}
