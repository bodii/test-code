package excel_sheet_column_title;

import java.util.ArrayList;
import java.util.Set;

class Solution {
    public String convertToTitle(int n) {
        StringBuilder result = new StringBuilder();
        while (n != 0) {
            result.insert(0, numberToChar(--n));
            n /= 26;
        }
        
        return result.toString();
    }

    public char numberToChar(int n) {
        return (char) ('A' + n % 26);
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        System.out.println(s.convertToTitle(701));
    }
}
