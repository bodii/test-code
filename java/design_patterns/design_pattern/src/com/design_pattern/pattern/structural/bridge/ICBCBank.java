package com.design_pattern.pattern.structural.bridge;

public class ICBCBank extends Bank {
    public ICBCBank(Account account) {
        super(account);
    }

    @Override
    Account openAccount() {
        System.out.println("打开一个工商银行账号");
        account.openAccount();
        return account;
    }
}
