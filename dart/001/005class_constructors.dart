import 'dart:math';

const double xOrigin = 0;
const double yOrigin = 0;

class Point {
  double x = 0;
  double y = 0;
  double distanceFromOrigin = 0;

  // named constructor
  Point.origin()
      : x = xOrigin,
        y = yOrigin;

  Point.fromJson(Map<String, double> json)
      : x = json['x']!,
        y = json['y']! {
    print('In Point.fromJson(): ($x, $y)');
  }

  Point.withAssert(this.x, this.y) : assert(x >= 0) {
    print('In Point.withAssert(): ($x, $y)');
  }

  Point(double x, double y)
      : x = x,
        y = y,
        distanceFromOrigin = sqrt(x * x + y * y);
}

class Person {
  String? firstName;

  Person.fromJson(Map data) {
    print('in Person');
  }
}

class Employee extends Person {
  Employee.fromJson(Map data) : super.fromJson(data) {
    print('in Employee');
  }
}

void main() {
  var employee = Employee.fromJson({});
  print(employee);

  var withAssert = Point.withAssert(1.0, 2.0);
  print(withAssert);

  var p = Point2(2, 3);
  print(p.distanceFromOrigin);
}

class Point2 {
  final double x;
  final double y;
  final double distanceFromOrigin;

  Point2(double x, double y)
      : x = x,
        y = y,
        distanceFromOrigin = sqrt(x * x + y * y);
}
