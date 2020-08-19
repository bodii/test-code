package com.design_pattern.pattern.behavioral.strategy;

public class ManjianPromotionStrategy implements PromotionStrategy {
    @Override
    public void doPromotions() {
        System.out.println("满减促销，课程的价格大于指定价格时，直接在总价格上减去优惠价格");
    }
}
