import 'string_apis.dart';

void main() {
  print('42'.padLeft(5));
  print('42'.parseInt());

  try {
    dynamic d = '2';
    print(d.parseInt());
  } on NoSuchMethodError catch (e) {}

  var v = '2';
  print(v.parseInt());
}
