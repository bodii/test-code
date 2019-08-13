/**
 * This program demonstrates parameter passing in java.
 * @version 1.01 2019-08-13
 * @author wong
 */

public class ParamTest {
	public static void main(String[] args) {
		/*
		 * Test 1: Methods can't modify numeric parameters
		 */
		System.out.println("Testing tripleValue: ");
		double percent = 10;
		System.out.println("Before: percent=" + percent);
		tripleValue(percent);
		System.out.println("After: percent=" + percent);

		/*
		 * Test 2: Mehtod can change the state of object parameters
		 */
		System.our.println("\nTesting tripleSalary: ");
		Employee harry = new Employee("Harry", 50000);

		System.out.println("Beffer: salay=" + harry.getSalary());
		

	}

	public  static void tripleValue(double x) {
		x = 3 * x;
		System.out.println("End of method: x= " + x);
	}

	public static void tripleSalary(Employee x) {
		x.raiseSalary(200);
		System.out.println("End of method: salary=" + x.getSalary());
	}

	public static void swap(Employee x, Employee y) {
		Employee temp = x;
		x = y;
		y = temp;

		System.out.println("End of method: x=" + x.getName());
		System.out.println("End of method: y=" + y.getName());
	}
}

/**
 *  Employee class
 * 
 * @version 1.1 2019-08-13
 * @author wong
 */
class Employee {
	private String name;
	private double salary;

	public Employee(String name, double salary) {
		this.name = name;
		this.salary = salary;
	}

	public String getName() {
		return name;
	}

	public double getSalary() {
		return salary;
	}

	public void raiseSalary(double byPercont) {
		double raise = salary * byPercont / 100;
		salary += raise;
	}
}