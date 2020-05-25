package chapter01;

public class Bank {
    private int money;
    private String name;

    public Bank(String name, int money) {
        this.money = money;
        this.name = name;
    }

    // 存款（deposit）
    public synchronized void deposit(int m) {
        money += m;
    }

    // 取款
    public synchronized boolean withdraw(int m) {
        if (money >= m) {
            money -= m;
            return true; 
        } 

        return false;
    }

    public String getName() {
        return name;
    }
}