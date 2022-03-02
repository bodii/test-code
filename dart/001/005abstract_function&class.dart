abstract class Doer {
  void doSomething();
}

class EffectiveDoer extends Doer {
  void doSomething() {}
}

abstract class AbstractContainer {
  void updateChildren();
}
