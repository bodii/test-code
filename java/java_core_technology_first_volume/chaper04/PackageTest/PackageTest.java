import com.horstmann.corejava.*;
// the Employee class is defined in that package

import static java.lang.System.*;

/**
 * This program demonstrates the use of packages.
 * @version 1.1 2019-08-13
 * @author wong
 */

 public class PackageTest {
     public static void main(String[] args) {
         Employee harry = new Employee("Harry Hacker", 50000, 2019, 8, 25);

         harry.raiseSalary(5);

         out.println("name=" + harry.getName());
     }
 }