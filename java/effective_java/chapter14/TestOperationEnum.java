package chapter14;

enum OperationEnum {
    PLUS {
        // 实例中实现抽象方法
        public double apply(double x, double y) {
            return x + y;
        }
    },
    MINUS {
        public double apply(double x , double y) {
            return x - y;
        }
    };

    // 声明抽象方法 
    public abstract double apply(double x, double y);
}


public class TestOperationEnum {
    public static void main(String[] args) {
        double result = OperationEnum.PLUS.apply(1, 2);
        System.out.println(result);
    }
}
