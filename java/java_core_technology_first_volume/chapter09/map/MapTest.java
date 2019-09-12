package map;

import java.util.*;

/**
 * This program demonstrates the use of a map of a map with key type String and value
 * type Employee.
 *
 * @version 1.1 2019-09
 * @author wong
 */
public class MapTest {
	public static void main(String[] args) {
		Map<String, Employee> staff = new HashMap<>();
		staff.put("144-25-5464" , new Employee("Amy lee"));
		staff.put("567-24-2546" , new Employee("Harry Hacker"));
		staff.put("157-62-7935" , new Employee("Gary Cooper"));
		staff.put("456-62-5527" , new Employee("Francesca Cruz"));

		System.out.println(staff);

		staff.remove("567-24-2546");

		staff.put("456-62-5227", new Employee("Francesca Miller"));

		System.out.println(staff.get("157-62-7935"));

		staff.forEach((k, v) -> 
			System.out.println("key=" + k + ", value=" + v));
	}
}

class Employee {
	private String name;

	Employee(String name) {
		this.name = name;
	}

	public String getName() {
		return name;
	}
}
