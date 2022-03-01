void main() {
  var halogens = {'fluorine', 'chlorine', 'bromine', 'iodine', 'astatine'};

  Set<String> halogens2 = {
    'fluorine',
    'chlorine',
    'bromine',
    'iodine',
    'astatine'
  };

  var names = <String>{};

  Set<String> names2 = {};
  var names3 = {}; // is a map

  var elements = <String>{};
  elements.add('fluorine');
  elements.addAll(halogens);

  assert(elements.length == 5);

  final constantSet = const {
    'fluorine',
    'chlorine',
    'bromine',
    'iodine',
    'astatine',
  };

  // constantSet.add('helium') // X
}
