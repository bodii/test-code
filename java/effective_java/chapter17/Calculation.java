package chapter17;

/**
 * 计算货币时的错误解答
 * 应该考虑使用BigDecimal、int或者long进行货币计算。
 */
public class Calculation {
    public static void main(String[] args) {
        double funds = 1.00;
        int itemsBought = 0;
        for (double price = 0.10; funds >= price; price += 0.10) {
            funds -= price;
            itemsBought ++;
        }
        System.out.println(itemsBought + " items bought.");
        System.out.println("Change: $" + funds);
    }
}