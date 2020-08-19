package com.design_pattern.pattern.behavioral.strategy;

import java.util.Map;

public class PromotionStrategyFactory {
    private static Map<String, PromotionStrategy> PROMOTION_STRATEGY = null;
    static {
        PROMOTION_STRATEGY.put(PromotionKey.FANXIAN, new FanxianPromotionStrategy());
        PROMOTION_STRATEGY.put(PromotionKey.LIJIAN, new LijianPromotionStrategy());
        PROMOTION_STRATEGY.put(PromotionKey.MANJIAN, new ManjianPromotionStrategy());
    }

    private static final PromotionStrategy EMPTY_PROMOTION_STRATEGY = new EmptyPromotionStrategy();

    private PromotionStrategyFactory() {}

    public static PromotionStrategy getPromotionStrategy(String promotionKey) {
        PromotionStrategy promotionStrategy = PROMOTION_STRATEGY.get(promotionKey);
        return promotionStrategy != null ? promotionStrategy: EMPTY_PROMOTION_STRATEGY;
    }

    private interface PromotionKey {
        final String FANXIAN = "FANXIAN";
        final String LIJIAN = "LIJIAN";
        final String MANJIAN = "MANJIAN";
    }
}
