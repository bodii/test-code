package com.design_pattern.pattern.structural.proxy;

public class OrderServiceImpl implements IOrderService {
    private IOrderDao orderDao;

    @Override
    public int saveOrder(Order order) {
        orderDao = new OrderDaoImpl();
        System.out.println("Service层调用Dao层添加Order");

        return orderDao.insert(order);
    }
}
