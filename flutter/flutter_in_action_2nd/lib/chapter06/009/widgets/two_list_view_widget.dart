import 'package:flutter/material.dart';

class TwoListViewWidget extends StatelessWidget {
  const TwoListViewWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return buildTwoListView();
  }
}

Widget buildTwoListView() {
  var listView1 = ListView.builder(
    itemCount: 20,
    itemBuilder: (_, index) => ListTile(title: Text('$index')),
    controller: ScrollController(),
  );
  var listView2 = ListView.builder(
    itemCount: 20,
    itemBuilder: (_, index) => ListTile(title: Text('$index')),
    controller: ScrollController(),
  );

  return Column(
    children: [
      Expanded(child: listView1),
      const Divider(color: Colors.grey),
      Expanded(child: listView2),
    ],
  );
}
