package com.design_pattern.pattern.structural.flyweight;

import java.util.HashMap;
import java.util.Map;

public class EmployeeFactory {
    private static final Map<String, Employee> EMPLOYEE_MAP = new HashMap<String, Employee>();

    public static Employee getManager(String departement) {
        Manager manager = (Manager) EMPLOYEE_MAP.get(departement);

        if (manager == null) {
            manager = new Manager(departement);
            System.out.println("创建部门经理： " + departement);
            String reportContent = departement + " 部门汇报：此次报告的主要内容是......";
            manager.setReportContent(reportContent);
            System.out.println("创建报告: " + reportContent);
            EMPLOYEE_MAP.put(departement, manager);
        }

        return manager;
    }
}
