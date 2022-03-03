class Foo {
  foo(x) {}
}

class MockFoo implements Foo {
  _mockFoo(x) {}

  @override
  noSuchMethod(Invocation invocation) {
    if (invocation.memberName == #foo) {
      if (invocation.isMethod &&
          invocation.positionalArguments.length == 1 &&
          invocation.namedArguments.isEmpty) {
        return _mockFoo(invocation.positionalArguments[0]);
      } else if (invocation.isGetter) {
        return _mockFoo;
      }
    }
    return super.noSuchMethod(invocation);
  }
}

class A {
  dynamic noSuchMethod(Invocation i) => null;

  void foo();
}

main() {
  A a = new A();
  dynamic f = a.foo;
  f();
}
