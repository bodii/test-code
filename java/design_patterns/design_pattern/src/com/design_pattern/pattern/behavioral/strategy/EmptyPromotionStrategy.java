package com.design_pattern.pattern.behavioral.strategy;

public class EmptyPromotionStrategy implements PromotionStrategy {

    @Override
    public void doPromotions() {
        System.out.println("无促销活动");
    }
}
