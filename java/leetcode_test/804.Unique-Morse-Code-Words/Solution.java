package unique_morse_code_words;

import java.util.TreeSet;

class Solution {
    public int uniqueMorseRepresentations(String[] words) {
        TreeSet<String> set = new TreeSet<>();
        String[] morses = {".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."};

        for (String word : words) {
            StringBuilder codeStr = new StringBuilder();
            for (int i = 0; i < word.length(); i++)
                codeStr.append(morses[word.charAt(i) - 'a']);
                
            set.add(codeStr.toString());
        }

        return set.size();
    }
}

class Test {
    public static void main(String[] args) {
        String[] words = {"gin", "zen", "gig", "msg"};

        Solution solution = new Solution();
        System.out.println(solution.uniqueMorseRepresentations(words));
    }
}