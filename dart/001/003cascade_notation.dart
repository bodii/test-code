void main() {
  var paint = Paint()
    ..color = Colors.black
    ..strokeCap = StrokeCap.round
    ..strokeWidth = 5.0;

  querySelector('#confirm')
    ?..text = 'Confirm'
    ..classes.add('important')
    ..onClick.listen((e) => window.alert('confirmed!'));
}

class Paint {
  late final String color;
  late final int strokeCap;
  late final double strokeWidth;
}

class Colors {
  static String black = '';
}

class StrokeCap {
  static int round = 0;
}

Button querySelector(String name) {
  return Button(name);
}

class Button {
  String name = '';
  String? text;
  Classes classes = Classes();
  Event onClick = Event();

  Button(this.name);
}

class Classes {
  void add(String? name) {}
}

class Event {
  void listen(Function func) {}
}

class window {
  static void alert(String? content) {}
}
