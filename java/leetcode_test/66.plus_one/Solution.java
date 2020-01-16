package plus_one;

class Solution {
    public int[] plusOne(int[] digits) {
        int len = digits.length;
        if (digits[len - 1] < 9 ) return digits;


        return digits;
    }
}

class Test {
    public static void main(String[] args) {
        Solution s = new Solution();

        int[] digits = new int[]{1, 2, 3};

        int[] newDigits = s.plusOne(digits);
        System.out.println("加一后的数组值是：");
        int len = newDigits.length;
        for (int i = 0; i < len; i++) {
            System.out.print(newDigits[i]);
            if (i != len - 1) 
                System.out.print(", ");
        }

        System.out.println();
    }
}