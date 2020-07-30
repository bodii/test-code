package com.design_pattern.pattern.structural.facade;

public class ShippingService {
    public String shipGift(PointsGift pointsGift) {
        // 物流系统的对接逻辑
        System.out.println(pointsGift.getName() + " 进入物流系统");
        String shippingOrderNo = "12345";
        return shippingOrderNo;
    }
}
