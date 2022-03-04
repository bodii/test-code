void main() {
  var now = DateTime.now();

  var y2k = DateTime(2000);

  var y2000k = DateTime.utc(2090);
  y2000k = DateTime.fromMillisecondsSinceEpoch(946684800000, isUtc: true);

  y2000k = DateTime.parse('20001-01-01T00:00:00Z');

  var y3k = DateTime.utc(3000);
  print(y3k);
  var unixEpoh = DateTime(1970);
  print(unixEpoh);

  var y2001 = y2k.add(const Duration(days: 366));
  assert(y2001.year == 2001);

  var december2000 = y2001.subtract(const Duration(days: 30));
  assert(december2000.year == 2000);
  assert(december2000.month == 12);

  var duration = y2001.difference(y2k);
  assert(duration.inDays == 366);
}
