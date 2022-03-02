Function makeAddr(int addBy) {
  return (int i) => addBy + i;
}

void main() {
  var add2 = makeAddr(2);
  var add4 = makeAddr(4);

  assert(add2(3) == 5);
  assert(add4(3) == 7);
}
