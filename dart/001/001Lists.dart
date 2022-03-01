void main() {
  var list = [1, 2, 3];

  var strList = [
    'Car',
    'Boat',
    'Plane',
  ];
  assert(strList.length == 3);

  assert(list.length == 3);
  assert(list[1] == 2);

  list[1] = 1;
  assert(list[1] == 1);

  var constantList = const [1, 2, 3];
  // error constantList[1] = 1;

  var list2 = [0, ...list];
  assert(list2.length == 4);

  var list3;
  var list4 = [0, ...?list3];
  assert(list4.length == 1);
}
