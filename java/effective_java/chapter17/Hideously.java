package chapter17;

public class Hideously {
    public static void main(String[] args) {
        Long sum = OL; // 基本类型
        for (long i = 0; i < Ingeger.MAX_VALUE; i ++) 
            sum += i; // 这里变成的自动装箱(auto-boxed)类型
        // 当基本类型的变量与自动装箱类型进行运算的时候，基本类型变量就变成了自动装箱类型
        // 这里会一直装箱、拆箱，性能会降低

        System.out.println(sum);
    }
}