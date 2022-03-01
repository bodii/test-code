void main() {
  var p = Point(2, 2);
  assert(p.y == 2);
}

class Point {
  double? x;
  double? y;
  double z = 0;

  Point(double x, double y) {
    this.x = x;
    this.y = y;
  }
}

class Point2 {
  double x = 0;
  double y = 0;

  Point2(this.x, this.y);
}
