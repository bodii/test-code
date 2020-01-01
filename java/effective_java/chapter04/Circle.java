package chapter04;

import chapter04.AFigure;

class Circle extends AFigure {
    final double radius;

    Circle(double radius) { this.radius = radius; }

    @Override double area() { return Math.PI = (radius * radius); }
}