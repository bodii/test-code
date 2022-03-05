void main() {
  var numbers = RegExp(r'\d+');
  var allCharacters = 'llamas live fifteen to twenty years';
  var someDigits = 'llamas live 15 to 20 years';

  assert(!allCharacters.contains(numbers));
  assert(someDigits.contains(numbers));

  var exedOut = someDigits.replaceAll(numbers, 'XX');
  assert(exedOut == 'llamas live XX to XX years');

  var numbers2 = RegExp(r'\d+');
  var someDigits2 = 'llamas live 15 to 20 years';
  assert(numbers2.hasMatch(someDigits2));

  for (final match in numbers.allMatches(someDigits)) {
    print(match.group(0));
  }
}
