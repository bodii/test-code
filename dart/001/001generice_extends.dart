class Foo<T extends Object> {}

class Foo2<T extends SomBaseClass> {
  String toString() => "Instance of 'Foo2<$T>'";
}

class Extender extends SomBaseClass {}

class SomBaseClass {}

void main() {
  var someBaseClassFoo = Foo<SomBaseClass>();
  var extenderFoo = Foo2<SomBaseClass>();

  var foo = Foo();
  print(foo);
}
