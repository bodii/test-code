package com.design_pattern.pattern.structural.flyweight;

public class Test {
    private static final String departements[] = {"RD", "QA", "PM", "BD"};

    public static void main(String[] args) {
        for (int i = 0; i < 10; i++) {
            String departement = departements[(int)(Math.random() * departements.length)];
            Manager manager = (Manager) EmployeeFactory.getManager(departement);
            manager.report();
        }

    }
}
