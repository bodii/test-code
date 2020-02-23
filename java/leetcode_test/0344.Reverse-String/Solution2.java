package reverse_string;


/*
    解题思路:
    Ideas for solving this problem:

    对字符数组进行折半，求得中间数(length/2), 遍历中间数长度，算出左右索引位，进行前后位的
    交换。
    当数组长度是奇数时，中间位不用考虑交换。
*/
class Solution2 {
    public void reverseString(char[] s) {
        int mid = s.length / 2;
        for (int left = 0; left < mid; left ++) {
            int right = s.length - 1 - left;
            
            char temp = s[left];
            s[left] = s[right];
            s[right] = temp;
        }
    }
}


class Test2 {
    public static void main(String[] args) {
        Solution2 solution = new Solution2();

        char[] test1 = {'h','e','l','l','o'};
        solution.reverseString(test1);
        for (char c : test1)
            System.out.print( c + ", ");

        System.out.println();

        char[] test2 = {'H','a','n','n','a','h'};
        solution.reverseString(test2);
        for (char c : test2)
            System.out.print( c + ", ");

        System.out.println();

        char[] test3 = {'A',' ','m','a','n',',',' ','a',' ','p','l','a','n',',',' ','a',' ','c','a','n','a','l',':',' ','P','a','n','a','m','a'};
        solution.reverseString(test3);
        System.out.println(test3.length);
        for (char c : test3)
            System.out.print( "\"" + c + "\", ");

        System.out.println();


    }
}