package com.design_pattern.pattern.behavioral.strategy;

public class Test {
    public static void main(String[] args) {
        PromotionActivity promotionstrategy618 = new PromotionActivity(new LijianPromotionStrategy());
        PromotionActivity promotionStrategy1111 = new PromotionActivity(new ManjianPromotionStrategy());

        promotionstrategy618.executePromotion();
        promotionStrategy1111.executePromotion();
    }
}
