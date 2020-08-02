package com.design_pattern.pattern.structural.bridge;

public class ABCBank extends Bank {
    public ABCBank(Account account) {
        super(account);
    }

    @Override
    Account openAccount() {
        System.out.println("打开一个农业银行账号");
        account.openAccount();
        return account;
    }
}
