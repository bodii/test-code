import 'dart:async';
import 'dart:math';
import 'dart:io';

void main() {
  var charCodes = [68, 97, 114, 116];
  var buffer = StringBuffer();

  charCodes.forEach(print);
  charCodes.forEach(buffer.write);

  var strings = charCodes.map(String.fromCharCode);
  var buffers = charCodes.map(StringBuffer.new);

  var lists = <num>[1, 2];
  lists.addAll(List<num>.filled(3, 4));
  lists.cast<int>();
}

int median(List<Object> objects) {
  var ints = List<int>.from(objects);
  ints.sort();
  return ints[ints.length ~/ 2];
}

void insert(Object item, {int at = 0}) {}

class Circle {
  double radius;
  Circle(this.radius);

  double get area => pi * radius * radius;
  double get circumference => pi * 2.0 * radius;
}

class Box {
  Object? value;

  void clear() {
    update(null);
  }

  void update(Object? value) {
    this.value = value;
  }
}

class ShadeOfGray {
  final int brightness;
  ShadeOfGray(int val) : brightness = val;
  ShadeOfGray.black() : this(0);
  ShadeOfGray.alsoBlack() : this.black();
}

var d = ShadeOfGray.black();

abstract class BaseBox {
  BaseBox(Object? value);
}

class Box2 extends BaseBox {
  Object? value;

  Box2(Object? value)
      : value = value,
        super(value);
}

class ProfileMark {
  final String name;
  final DateTime start = DateTime.now();

  ProfileMark(this.name);
  ProfileMark.unnamed() : name = '';
}

Future<bool> fileContainsBear(String path) {
  return File(path).readAsString().then((contents) {
    return contents.contains('bear');
  });
}

Future<bool> fileContainsBear2(String path) async {
  var contents = await File(path).readAsString();
  return contents.contains('bear');
}

///
/// # aaa
/// ```dart
/// print(a)
/// ```
///
Future<T> logValue<T>(FutureOr<T> value) async {
  if (value is Future<T>) {
    var result = await value;
    print(result);
    return result;
  } else {
    print(value);
    return value;
  }
}
