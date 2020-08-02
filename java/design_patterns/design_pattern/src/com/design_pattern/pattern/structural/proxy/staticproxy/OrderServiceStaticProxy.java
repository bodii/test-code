package com.design_pattern.pattern.structural.proxy.staticproxy;

import com.design_pattern.pattern.structural.proxy.IOrderService;
import com.design_pattern.pattern.structural.proxy.Order;
import com.design_pattern.pattern.structural.proxy.OrderServiceImpl;

public class OrderServiceStaticProxy {
    private IOrderService iOrderService;

    public int saveOrder(Order order) {
        beforeMethod();
        iOrderService = new OrderServiceImpl();

        int userId = order.getUserId();
        int dbRouter = userId % 2;
        System.out.println("静态代理分配到 [db " + dbRouter + "] 处理数据");

        afterMethod();
        return iOrderService.saveOrder(order);
    }

    private void beforeMethod() {
        System.out.println("静态代理 before code");
    }

    private void afterMethod() {
        System.out.println("静态代理 after code");
    }
}
