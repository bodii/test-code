package unique_morse_code_words;

import java.util.Set;
import java.util.HashSet;

class Solution2 {
    private final String[] morses = new String[]{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."};

    public int uniqueMorseRepresentations(String[] words) {
        Set<String> set = new HashSet<>();
        
        for (String word : words) {
            StringBuilder codeStr = new StringBuilder();
            for (char c : word.toCharArray())
                codeStr.append(morses[c - 97]);
                
            set.add(codeStr.toString());
        }

        return set.size();
    }
}

class Test2 {
    public static void main(String[] args) {
        String[] words = {"gin", "zen", "gig", "msg"};

        Solution2 solution = new Solution2();
        System.out.println(solution.uniqueMorseRepresentations(words));
    }
}