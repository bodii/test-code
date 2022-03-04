var teas = ['green', 'black', 'chamomile', 'earl grey'];

void main() {
  var loudTeas = teas.map((tea) => tea.toUpperCase()).toList();
  var loudTeas2 = teas.map((tea) => tea.toLowerCase()).toSet();

  print(loudTeas);
  print(loudTeas2);
}
