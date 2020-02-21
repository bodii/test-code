package plus_one;

class Solution {
    public int[] plusOne(int[] digits) {
        digits[digits.length -1] ++;
        if (digits[digits.length -1] != 10)
            return digits;

        for (int i = digits.length - 1; i >= 0; i --) {
            if (digits[i] == 10) {
                digits[i] = 0;
                if (i > 0)
                    digits[i -1] ++;
            } else 
                return digits;

        }

        digits = new int[digits.length + 1];
        digits[0] = 1;

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


        int[] digits2 = new int[]{4,3,2,1};

        int[] newDigits2 = s.plusOne(digits2);
        System.out.println("加一后的数组值是：");
        int len2 = newDigits2.length;
        for (int i = 0; i < len2; i++) {
            System.out.print(newDigits2[i]);
            if (i != len2 - 1) 
                System.out.print(", ");
        }

        System.out.println();


        int[] digits3 = new int[]{9,8,9};

        int[] newDigits3 = s.plusOne(digits3);
        System.out.println("加一后的数组值是：");
        int len3 = newDigits3.length;
        for (int i = 0; i < len3; i++) {
            System.out.print(newDigits3[i]);
            if (i != len3 - 1) 
                System.out.print(", ");
        }

        System.out.println();
    }
}