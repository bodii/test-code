package arrayList;

import java.util.*;

import arrayList.Employee;

/**
 * This program demonstrates the ArrayList class.
 * @version .1.1 2019-08
 * @author wong
 */

public class ArrayListTest {
    public static void main(String[] args) {
        ArrayList<Employee> staff = new ArrayList<>();

        staff.add(new Employee("Carl Cracker", 75000, 1987, 12, 15));
        staff.add(new Employee("Harry Hacker", 50000, 1989, 10, 1));
        staff.add(new Employee("Tony Tester", 40000, 1990, 3, 15));

        for (Employee e : staff)
            e.raiseSalary(5);

        for (Employee e : staff)
            System.out.printf(
                "name = %s, salary = %g, hireDay = %s%n", 
                e.getName(), e.getSalary(), e.getHireDay()
            );
    }
}