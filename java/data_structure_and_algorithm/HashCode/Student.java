package hash_code;

import java.util.HashMap;
import java.util.HashSet;

public class Student {
    int grade;
    int classes;
    String firstName;
    String lastName;

    public Student(int grade, int classes, String firstName, String lastName) {
        this.grade = grade;
        this.classes = classes;
        this.firstName = firstName;
        this.lastName = lastName;
    }

    @Override
    public int hashCode() {
        int B = 31;

        int hash = 0;
        hash = hash * B + grade;
        hash = hash * B + classes;
        // hash = hash * B + firstName.hashCode();
        // hash = hash * B + lastName.hashCode();
        // 不区分大小写
        hash = hash * B + firstName.toLowerCase().hashCode();  
        hash = hash * B + lastName.toLowerCase().hashCode();

        return hash;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o)
            return true;

        if (o == null) 
            return false;

        if (getClass() != o.getClass())
            return false;

        Student another = (Student) o;
        return this.grade == another.grade && this.classes == another.classes
                    && this.firstName.toLowerCase().equals(firstName.toLowerCase())
                    && this.lastName.toLowerCase().equals(another.lastName.toLowerCase());
    }

    public static void main(String[] args) {
        int a = 42;
        System.out.println(((Integer)a).hashCode());

        int b = -42;
        System.out.println(((Integer)b).hashCode());

        double c = 3.1415926;
        System.out.println(((Double)c).hashCode());

        String d = "abc";
        System.out.println(d.hashCode());

        Student student = new Student(3, 2, "ping", "wu");
        System.out.println(student.hashCode());

        HashSet<Student> set = new HashSet<>();
        set.add(student);

        HashMap<Student, Integer> map = new HashMap<>();
        map.put(student, 10);

        Student student2 = new Student(5, 6, "meimei", "han");
        System.out.println(student2.hashCode());
    }
}