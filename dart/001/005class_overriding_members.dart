class Television {
  set contrast(int value) {}
}

class SmartTelevision extends Television {
  @override
  set contrast(num value) {}
}

class Animal {
  void chase(Animal x) {}
}

class Mouse extends Animal {}

class Cat extends Animal {
  @override
  void chase(covariant Mouse x) {}
}
