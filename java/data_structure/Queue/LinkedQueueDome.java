import java.util.Scanner;

/**
 * 链式队列演示类-排队食堂打饭
 */
public class LinkedQueueDome {
    public static void main(String[] args) {
        LinkedQueue canteenQueue = new LinkedQueue(50);
        Scanner input = new Scanner(System.in);
        System.out.println("请输入前来打饭的人员姓名(f表示结束输入):");
        String name = input.next();

        while (!name.equals("f")) {
            canteenQueue.push(name); // 来一个人将他排入队伍
            name = input.next(); // 获取下一个打饭的人
        }

        while (!canteenQueue.isEmpty()) { // 如果食堂不为空
            System.out.println(canteenQueue.pop() + " 打完饭离开队伍。");
        }
    }
}