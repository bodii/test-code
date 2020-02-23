package reverse_string;


/*
    解题思路:
    Ideas for solving this problem:

*/
class Solution {
    private void swap(char[]s, int l, int r) {
        if (l >= r)
            return;

        char temp = s[l];
        s[l++] = s[r];
        s[r--] = temp;

        swap(s, l, r);
    }

    public void reverseString(char[] s) {
        swap(s, 0, s.length - 1);
    }
}


class Test {
    public static void main(String[] args) {
        Solution solution = new Solution();

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
        for (char c : test3)
            System.out.print( "\"" + c + "\", ");

        System.out.println();


    }
}