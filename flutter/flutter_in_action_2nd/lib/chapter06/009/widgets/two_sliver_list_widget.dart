import 'package:flutter/material.dart';

class TwoSliverListWidget extends StatelessWidget {
  const TwoSliverListWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return buildTwoSliverList();
  }
}

Widget buildTwoSliverList() {
  var listView = SliverFixedExtentList(
    itemExtent: 56, // 列表项高度固定
    delegate: SliverChildBuilderDelegate(
      (_, index) => ListTile(title: Text('$index')),
      childCount: 10,
    ),
  );

  return CustomScrollView(
    slivers: [
      listView,
      listView,
    ],
  );
}
