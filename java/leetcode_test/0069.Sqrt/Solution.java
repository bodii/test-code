package sqrt;

/*
    解题思路：
    牛顿迭代法。就是不断取切线，用线性方程的根逼近非线性方程f(x)=0的根X*。
*/
class Solution {
    public int mySqrt(int x) {
        double pre = x, t = (pre + x / pre) / 2;

        while (Math.abs(t * t - x) > 0.000001) 
            t= (t + x / t) / 2;

	    return (int)t;
    }
}


class Test {
    public static void main(String[] args) {
        Solution solution = new Solution();
        int x = 4;
        System.out.println( x + " sqrt is : " + solution.mySqrt(x));

        x = 8;
        System.out.println( x + " sqrt is : " + solution.mySqrt(x));

        x = 33;
        System.out.println( x + " sqrt is : " + solution.mySqrt(x));
    }
}