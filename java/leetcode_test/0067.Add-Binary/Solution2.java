package add_binary;

class Solution2 {
    public String addBinary(String a, String b) {
        int inta = Integer.parseInt(a, 2);
        int intb = Integer.parseInt(b, 2);
    
        return Integer.toBinaryString(inta + intb);
    }

    public static void main(String[] args) {
        Solution2 s = new Solution2();
        String a = "11";
        String b = "1";
        System.out.println(s.addBinary(a, b));
    }
}
