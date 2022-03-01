void main() {
  bool promoActive = true;
  var nav = ['Home', 'Furniture', 'Plats', if (promoActive) 'Outlet'];

  print(nav);

  var listOfInts = [1, 2, 3];
  var listOfStrings = ['#0', for (int i in listOfInts) '#$i'];
  print(listOfStrings);
  assert(listOfStrings[1] == '#1');
}
