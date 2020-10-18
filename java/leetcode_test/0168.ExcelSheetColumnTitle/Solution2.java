package excel_sheet_column_title;

class Solution2 {
    public String convertToTitle(int n) {
        String result = "";
        while (n != 0) {
            result = String.valueOf(numberToChar(--n)) + result;
            n /= 26;
        }

        return result;
    }

    public char numberToChar(int n) {
        return (char) ('A' + n % 26);
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        System.out.println(s.convertToTitle(701));
    }
}
