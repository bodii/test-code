package com.design_pattern.pattern.behavioral.strategy;

public class PromotionActivity {
	private PromotionStrategy promotionStrategy;

	public PromotionActivity(PromotionStrategy promotionStrategy) {
		this.promotionStrategy = promotionStrategy;
	}

	public void executePromotion() {
		promotionStrategy.doPromotions();
	}
}
