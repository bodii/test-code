package enums;

import java.util.*;

/**
 * This program demonstrates enumerated types
 * @version 1.1
 * @author wong
 */

 public class  EnumTest {
    public static void main(String[] args) {
        Scanner stdIn = new Scanner(System.in);
        System.out.print("Enter a size: (SMALL, MEDIUM, LARGE, EXTRA_LARGE) >> ");
        String input = stdIn.next().toUpperCase();
        Size size = Enum.valueOf(Size.class, input);
        System.out.println("size = " + size);
        System.out.println("abbreviation = " + size.getAbbreviation());
        if (size == Size.EXTRA_LARGE)
            System.out.println("Good job--you paid attention to the _.");
    }
 }

 enum Size {
     SMALL("S"), MEDIUM("M"), LARGE("L"), EXTRA_LARGE("XL");

     private Size(String abbreviation) { this.abbreviation = abbreviation; }
     public String getAbbreviation() { return abbreviation; }

     private String abbreviation;
 }