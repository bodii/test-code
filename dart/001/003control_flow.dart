void main() {
  Object name = 'Bob';
  String name2 = 'Bob';
  // int? lineCount;
  // assert(lineCount != null, 'lineCount is null');

  print(name);
  print(name2);

  // try {
  //   throw "adfdsfds";
  // } on OutOfMemoryError {
  // } on Exception catch (e) {
  //   print('Unknown exception: $e');
  // } catch (e) {
  //   print('Something really unknown: $e');
  //   rethrow;
  // } finally {
  //   print('finally');
  // }

  var year = 2000;
  if (year >= 2001) {
    print('21st century');
  } else if (year >= 1901) {
    print('20th century');
  }

  const flybyObjects = ['Jupiter', 'Saturn', 'Uranus', 'Neptune'];
  for (final object in flybyObjects) {
    print(object);
  }

  for (int month = 1; month <= 12; month++) {
    print(month);
  }

  while (year < 2016) {
    year++;
  }
}
