package com.design_pattern.pattern.structural.decorator.v2;

public class Test {
    public static void main(String[] args) {
        ABattercake battercake = new Battercake();
        System.out.println(battercake.getDose() + " 销售价格： " + battercake.cost());

        ABattercake egg = new EggDecorator(battercake);
        System.out.println(egg.getDose() + " 销售价格： " + egg.cost());
        ABattercake sausage = new SausageDoractor(battercake);
        System.out.println(sausage.getDose() + " 销售价格： " + sausage.cost());

        ABattercake eggAndSausage = new SausageDoractor(new EggDecorator(battercake));
        System.out.println(eggAndSausage.getDose() + " 销售价格： " + eggAndSausage.cost());
        ABattercake eggAndSausage2 = new SausageDoractor(eggAndSausage);
        System.out.println(eggAndSausage2.getDose() + " 销售价格： " + eggAndSausage2.cost());
    }
}
