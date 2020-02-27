package map_sum;

import java.util.TreeMap;

class MapSum {
    /**
     * 内部节点类
     */
    private class Node {
        public int val;
        public boolean isWord;
        public TreeMap<Character, Node> next; 

        
        public Node(int val, boolean isWord) {
            this.val = val;
            this.isWord = isWord;
            next = new TreeMap<>();
        }

        public Node() {
            this(0, false);
        }
    } 

    private Node root;
    private int size;

    /** Initialize your data structure here. */
    public MapSum() {
        root = new Node();
        size = 0;
    }
    
    public void insert(String key, int val) {
        Node current = root;
        for (int i = 0; i < key.length(); i ++) {
            char c = key.charAt(i);
            if (current.next.get(c) == null)
                current.next.put(c, new Node());

            current = current.next.get(c);
        }

        current.val = val;
        if (!current.isWord) {
            current.isWord = true;
            size ++;
        }
    }
    
    public int sum(String prefix) {
        int result = 0;
        Node current = root;
        for(int i = 0; i < prefix.length(); i ++) {
            result += current.val;

            char c = prefix.charAt(i);
            if (current.next.get(c) == null)
                return result;
            current = current.next.get(c);
        }

        for (char nextChar : current.next.keySet()) {

        }
        return result + getval(current, c)
    }

    


    private int getVal(Node node, char nextChar) {
        if (node.next.get(nextChar) == null)
            return node.val;

        totalNum += node.val;


        return getVal(node.next.get(nextChar), c, totalNum);
    }

}

/**
 * Your MapSum object will be instantiated and called as such:
 * MapSum obj = new MapSum();
 * obj.insert(key,val);
 * int param_2 = obj.sum(prefix);
 */