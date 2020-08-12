package com.easykotlin.test.chapter12.reflectTest;


import java.lang.annotation.Annotation;
import java.lang.reflect.Field;
import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;
import java.util.Arrays;
import java.util.List;

interface StudentService<T> {
    List<T> findStudents(String name, Integer age);
}

abstract class BaseService<T> {
    abstract int save(T t);
}

class Student {
    String name;
    Integer age;

    public Student(String name, Integer age) {
        this.name = name;
        this.age = age;
    }

    public String  getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Integer getAge() {
        return age;
    }

    public void setAge(Integer age) {
        this.age = age;
    }
}

class StudentServiceImpl extends
        BaseService<Student> implements StudentService<Student> {
    @Override
    public List<Student> findStudents(String name, Integer age) {
        return Arrays.asList(new Student("Jack", 20),
                new Student("Rose", 20));
    }

    @Override
    public int save(Student student) {
        return 0;
    }
}


public class ReflectionDemo {
    public static void main(String[] args) {
        StudentServiceImpl studentService = new StudentServiceImpl();
        studentService.save(new Student("Bob", 20));
        studentService.findStudents("Jack", 20);

        // 反射API调用示例
        final Class<? extends StudentServiceImpl> studentServiceClass = studentService.getClass();
        Class<?>[] classes = studentServiceClass.getDeclaredClasses();
        Annotation[] annotations = studentServiceClass.getAnnotations();
        ClassLoader classLoader = studentServiceClass.getClassLoader();

        Field[] fields = studentServiceClass.getFields();
        Method[] methods = studentServiceClass.getMethods();

        try {
            System.out.println(methods[0].getName());
            methods[0].invoke(studentService, new Student("Jack", 20));
        } catch (IllegalAccessException | InvocationTargetException e) {
            e.printStackTrace();
        }
    }
}