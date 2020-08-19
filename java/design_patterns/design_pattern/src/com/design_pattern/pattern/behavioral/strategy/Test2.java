package com.design_pattern.pattern.behavioral.strategy;

public class Test2 {
    public static void main(String[] args) {
        PromotionActivity promotionStrategy = new PromotionActivity(PromotionStrategyFactory.getPromotionStrategy("LIJIAN"));
        promotionStrategy.executePromotion();
    }
}
