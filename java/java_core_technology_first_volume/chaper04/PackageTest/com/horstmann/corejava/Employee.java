package com.horstmann.corejava;
// the class in this file are part of this package

import java.time.*;
// import statements come after the package statement

/**
 * @version 1.1 2019-08-13
 * @author wong
 */

 public class Employee {
     private String name;
     private double salary;
     private LocalDate hireDay;

     public Employee(String name, double salary, int year, int month, int day) {
         this.name = name;
         this.salary = salary;
         hireDay = LocalDate.of(year, month, day);
     }

     public String getName() {
         return name;
     }

     public double getSalary() {
         return salary;
     }

     public LocalDate getHireDay() {
         return hireDay;
     }

     public void raiseSalary(double byPercent) {
         double raise = salary * byPercent / 100;
         salary += raise;
     }
 }