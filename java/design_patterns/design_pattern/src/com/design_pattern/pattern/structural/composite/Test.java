package com.design_pattern.pattern.structural.composite;

public class Test {
    public static void main(String[] args) {
        CatalogComponent  linuxCourse = new Course("Linux", 100);
        CatalogComponent windowsCourse = new Course("Windows", 40);
        CatalogComponent javaCourse = new Course("Java", 120);

        CatalogComponent java1Course = new Course("Java一期", 25);
        CatalogComponent java2Course = new Course("Java二期", 30);
        CatalogComponent java3Course = new Course("Java三期", 40);

        javaCourse.add(java1Course);
        javaCourse.add(java2Course);
        javaCourse.add(java3Course);

        CatalogComponent courseCatalog = new CourseCatalog("课程目录");
        courseCatalog.add(linuxCourse);
        courseCatalog.add(windowsCourse);
        courseCatalog.add(javaCourse);

        courseCatalog.print();
    }
}
