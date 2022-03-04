var teas = ['green', 'black', 'chamomile', 'earl grey'];
var hawaiianBeaches = {'waikicki': 1, 'kailua': 2, 'waimanalo': 3};

void main() {
  teas.forEach((tea) => print('I drink $tea'));

  hawaiianBeaches.forEach((k, v) {
    print('I want to visit $k and swim at $v');
  });

  var loudTeas = teas.map((tea) => tea.toUpperCase());
  loudTeas.forEach(print);
}
