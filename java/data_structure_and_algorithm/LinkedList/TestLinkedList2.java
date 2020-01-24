package linked_list;

/**
 * 链表类2 测试类
 */
public class TestLinkedList2 {
    public static void main(String[] args) {
        LinkedList2<Integer> linked = new LinkedList2<>();
        for (int i = 0; i < 5; i ++) {
            linked.addFirst(i);
            System.out.println(linked);
        }

        System.out.println("在链表中索引为的2，添加元素888:");
        linked.add(2, 888);
        System.out.println(linked);

        System.out.println("删除链表索引为的2的元素");
        linked.remove(2);
        System.out.println(linked);
    }
}