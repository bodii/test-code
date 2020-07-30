package com.design_pattern.pattern.decorator.v1;

public class Test {
    public static void main(String[] args) {
        Battercake battercake = new Battercake();
        System.out.println(battercake.getDesc() + " 销售价格：" + battercake.cost());

        Battercake battercakeWithEgg = new BattercakeWithEgg();
        System.out.println(
                battercakeWithEgg.getDesc()
                        + " 销售价格： "
                        + battercakeWithEgg.cost()
        );

        Battercake batercakeWithEggSausage = new BattercakeWithEggSausage();
        System.out.println(
                batercakeWithEggSausage.getDesc()
                + " 销售价格： "
                + batercakeWithEggSausage.cost()
        );

    }
}
