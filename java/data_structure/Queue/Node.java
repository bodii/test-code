/**
 * 节点类
 */
public class Node {
    private Object nodeValue;
    private Node next;

    public Node() {
        nodeValue = null;
        next = null;
    }

    public Node(Object item) {
        nodeValue = item;
        next = null;
    }

    public Node(Object item, Node node) {
        nodeValue = item;
        next = node;
    }

    public Object getNodeValue() {
        return nodeValue;
    }

    public void setNodeValue(Object item) {
        this.nodeValue = item;
    }

    public Node getNext() {
        return next;
    }

    public void setNext(Node next) {
        this.next = next;
    }
}