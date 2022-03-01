void main() {
  final name = 'Bob';
  final String nickname = 'Bobby';

  // name = 'Alice'; // X

  const bar = 1000000;
  const double atm = 1.01325 * bar;
  print(name);
  print(nickname);

  var foo = const [];
  final bar2 = const [];
  const baz = [];
  print(foo);
  print(bar2);
  print(baz);

  foo = [1, 2, 3];
  print(foo);
}
