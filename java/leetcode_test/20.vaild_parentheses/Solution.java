package vaild_parentheses;

import java.util.Stack;

class Solution {
    public boolean isValid(String s) {
        Stack<Character> stack = new Stack<>();
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (c == '(' || c =='[' || c == '{') {
                stack.push(c);
            } else {
                if (stack.isEmpty()) return false;

                char topChar = stack.pop();
                if (c == ')' && topChar != '(') return false;
                if (c == ']' && topChar !=  '[') return false;
                if (c == '}' && topChar !=  '{') return false;
            }
        }

        return stack.isEmpty();
    }
}


class Test {
    public static void main(String[] args) {
        Solution s = new Solution();
        String str = new String("()[]{}");
        System.out.println("原字符串：" + str);
        System.out.println("字符串是否: " + s.isValid(str));
    }
}