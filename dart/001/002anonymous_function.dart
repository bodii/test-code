const list = ['apples', 'bananas', 'oranges'];

void main() {
  list.forEach((item) => print('${list.indexOf(item)}: $item'));
}
