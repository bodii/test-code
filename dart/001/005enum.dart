enum Color { red, green, blue }

void main() {
  assert(Color.red.index == 0);
  assert(Color.green.index == 0);
  assert(Color.blue.index == 0);

  List<Color> colors = Color.values;
  assert(colors[2] == Color.blue);

  var aColor = Color.blue;
  switch (aColor) {
    case Color.red:
      print('Red as roses!');
      break;
    case Color.green:
      print('Green as grass!');
      break;
    default:
      print(aColor);
  }
}
