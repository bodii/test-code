package length_of_last_word;

/*

解题思路

从右向左遍历，并声明一个返回值(result)为0数，用来记录最后一个单词的长度。
返回值(result)是0时，当前字符是 ’ ’ 时，说明还没有触探到最后一个单词的边界，跳过此次遍历机会；
当返回值(result)非0时，如果当前字符是’ ’ 时，说明最后一个单词的长度已经获取完毕结束遍历，否则返回值(result)+1；
最终返回result。

 */
class Solution3 {
    public int lengthOfLastWord(String s) {
        int right = s.length();
        int result = 0;
        while (--right >= 0) {
            if (result == 0 && ' ' == s.charAt(right))
                continue;
            if (' ' == s.charAt(right))
                break;
            else
                result ++;
        }

        return result;
    }
}

class Test3 {
    public static void main(String[] args) {
        Solution3 solution = new Solution3();
        String s = "Hello World";
        System.out.println(solution.lengthOfLastWord(s));

        String s1 = "a";
        System.out.println(solution.lengthOfLastWord(s1));

        String s2 = " ";
        System.out.println(solution.lengthOfLastWord(s2));
    }
}