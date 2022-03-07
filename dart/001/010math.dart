import "dart:math";

void main() {
  assert(cos(pi) == -1.0);
  var degrees = 30;
  var radians = degrees * (pi / 180);
  var sinOf30degrees = sin(radians);

  assert((sinOf30degrees - 0.5).abs() < 0.01);

  assert(max(1, 1000) == 1000);
  assert(min(1, -1000) == -1000);

  print(e);
  print(pi);
  print(sqrt2);

  var random = Random();
  random.nextDouble();
  random.nextInt(10);

  var random2 = Random();
  print(random2.nextBool());
}
