Stream<int> asynchronousNaturalsTo(int n) async* {
  int k = 0;
  while (k < n) yield k++;
}

void main() {
  asynchronousNaturalsTo(4);
}
