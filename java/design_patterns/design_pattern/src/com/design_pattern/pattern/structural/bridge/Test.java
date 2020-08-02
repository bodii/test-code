package com.design_pattern.pattern.structural.bridge;

public class Test {
    public static void main(String[] args) {
        Bank abcBankDepositAccount = new ABCBank(new DepositAccount());
        abcBankDepositAccount.openAccount().showAccountType();
        // bank -> 打开（委托） account->显示账户
        System.out.println();

        Bank abcBankSavingAccount = new ABCBank(new SavingAccount());
        abcBankSavingAccount.openAccount().showAccountType();
        System.out.println();

        Bank icbcBankSavingAccount = new ICBCBank(new SavingAccount());
        icbcBankSavingAccount.openAccount().showAccountType();
    }
}
