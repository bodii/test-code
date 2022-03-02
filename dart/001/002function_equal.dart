void foo() {}

class A {
  static void bar() {}
  void baz() {}
}

void main() {
  Function x;

  x = foo;
  assert(foo == x);

  x = A.bar;
  assert(A.bar == x);

  var v = A();
  var w = A();
  var y = w;
  x = w.baz;
  assert(y.baz == x);

  assert(v.baz != w.baz);
}
