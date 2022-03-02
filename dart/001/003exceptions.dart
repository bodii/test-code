void main() {
  t();

  try {
    void distanceTo(Point other) => throw UnimplementedError();

    throw FormatException('Expected at least 1 section');
  } on FormatException catch (e) {
    print('exception: $e');
  } catch (e) {
    rethrow; // 重新抛出异常
  } finally {
    print('finally');
  }
}

Error t() {
  throw 'Out of llamas!';
}

class Point {}
