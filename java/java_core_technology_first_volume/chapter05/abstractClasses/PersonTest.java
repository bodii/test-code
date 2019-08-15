package abstractClasses;

import abstractClasses.*;

/**
 * This program demonstrates abstract classes.
 * @version 1.1 2019-08
 * @author wong
 */
public class PersonTest {
    public static void main(String[] args) {
        Person[] people = new Person[2];

        //fill the people array with student and Employee objects
        people[0] = new Employee("Harry hacker", 50000, 1989, 10, 1);
        people[1] = new Student("Maria Morris", "computer science");

        // print out names and descriptions of all Person objects
        for (Person p : people)
            System.out.printf("name = %s, description = %d", p.getName(), p.getDescription());
    }
}