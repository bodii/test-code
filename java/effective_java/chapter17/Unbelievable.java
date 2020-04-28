package chapter17;

public class Unbelievable {
    public static Integer i;

    public static void main(String[] args) {
        i = new Integer(42); // 自动装箱(auto-boxed)数据类型

        if (i == 42) // 自动拆箱（auto-unboxed)
            System.out.println("Unblievable");
    } 
}