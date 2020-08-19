package com.design_pattern.pattern.behavioral.strategy;

public class FanxianPromotionStrategy implements PromotionStrategy {
    @Override
    public void doPromotions() {
        System.out.println("返现促销，返回的金额存入用户余额");
    }
}
