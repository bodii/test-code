package com.easykotlin.test.chapter6;

/**
 * Java 工具类StringUtil的代码
 */
public class StringUtil {
    public static String firstChar(String str) {
        String first = "";
        if (str != null && str.length() > 0) {
            return first + str.charAt(0);
        }

        return first;
    }

    public static String lastChar(String str) {
        String last = "";
        if (str != null && str.length() > 0) {
            return last + str.charAt(str.length() -1);
        }

        return last;
    }

    public static void main(String[] args) {
        String str = "abc";
        System.out.println(firstChar(str));
        System.out.println(lastChar(str));
    }
}
