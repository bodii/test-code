class Performer {}

class Person {
  String name;

  Person(this.name);
}

class Musician extends Performer with Musical {}

class Maestro extends Person with Musical, Aggressive, Demented {
  Maestro(String maestroName) : super('') {
    name = maestroName;
    canConduct = true;
  }
}

mixin Musical {
  bool canPlayPiano = false;
  bool canCompose = false;
  bool canConduct = false;

  void entertainMe() {
    if (canPlayPiano) {
      print('Playing piano');
    } else if (canConduct) {
      print('Waving hands');
    } else {
      print('Humming to self');
    }
  }
}

mixin Aggressive {}

mixin Demented {}
