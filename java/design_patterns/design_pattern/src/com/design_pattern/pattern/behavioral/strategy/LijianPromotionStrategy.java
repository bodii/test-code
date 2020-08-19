package com.design_pattern.pattern.behavioral.strategy;

public class LijianPromotionStrategy implements PromotionStrategy {
    @Override
    public void doPromotions() {
        System.out.println("立减促销，课程的价格直接减去配置的价格");
    }
}
