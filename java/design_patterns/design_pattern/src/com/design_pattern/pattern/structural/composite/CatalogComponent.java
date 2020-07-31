package com.design_pattern.pattern.structural.composite;

public abstract class CatalogComponent {
    public void add(CatalogComponent catalogComponent) {
        throw new UnsupportedOperationException("不支持添加操作");
    }

    public void remove(CatalogComponent catalogComponent) {
        throw new UnsupportedOperationException("不支持删除操作");
    }

    public String getName(CatalogComponent catalogComponent) {
        throw new UnsupportedOperationException("不支持获取操作");
    }

    public double getPrice(CatalogComponent catalogComponent) {
        throw new UnsupportedOperationException("不支持获取价格");
    }

    public void print() {
        throw new UnsupportedOperationException("不支持打印操作");
    }
}
