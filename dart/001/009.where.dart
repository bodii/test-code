var teas = ['green', 'black', 'chamomile', 'earl grey'];

bool isDecaffeinated(String teaName) => teaName == 'chamomile';

var decaffeinatedTeas = teas.where((tea) => isDecaffeinated(tea));

void main() {
  assert(teas.any(isDecaffeinated));
  assert(!teas.every(isDecaffeinated));
}
