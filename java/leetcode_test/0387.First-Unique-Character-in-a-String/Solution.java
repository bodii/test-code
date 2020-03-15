
class Solution {
    public int firstUniqChar(String s) {
		int[] words = new int[26];
		for (int i = 0; i < s.length(); i ++)
			words[s.charAt(i) - 'a'] ++;

		for (int i = 0; i < s.length(); i ++)
			if (words[s.charAt(i) - 'a'] == 1)
				return i;

		return -1;
    }
}

class Test {
	public static void main(String[] args) {
		Solution solution = new Solution();

		String str1 = "lettcode";
		System.out.println(solution.firstUniqChar(str1));

		String str2 = "lovelettcode";
		System.out.println(solution.firstUniqChar(str2));
	}
}
