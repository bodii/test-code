void main() {
  var names = <String>['Seth', 'Kathy', 'Lars'];
  var nameSet = Set<String>.from(names);

  var views = Map<int, View>();

  var names2 = <String>[];
  names2.addAll(['Seth', 'Kathy', 'Lars']);
  print(names is List<String>);
}

class View {}
